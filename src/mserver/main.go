/**
 * gomserver main.go
 */
package main

import (
	"base"
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	//"os/exec"
	//"handle"
	"io"
	"log"
	"net"
	"path/filepath"
)

const (
	TGW_HEADER_SIZE      = 1024 * 4
	TGW_HEADER_SEG_COUNT = 3
)

func main() {
	//守护进程，开始
	d := flag.Bool("d", false, "Whether or not to launch in the background(like a daemon)")
	flag.Parse()
	fmt.Println(*d)
	if *d {
		fmt.Println(os.Args[0] + " will run in background.")
		filePath, _ := filepath.Abs(os.Args[0]) //将命令行参数中执行文件路径转换成可用路径
		//cmd := exec.Command(filePath, os.Args[2:]...)
		//将其他命令传入生成出的进程
		//cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//cmd.Start() //开始执行新进程，不等待新进程退出
		args := append([]string{filePath}, os.Args[2:]...)
		os.StartProcess(filePath, args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		return
	}
	//守护进程，结束

	defer base.Defer()

	base.SayHello("gomserver is running.")
	base.SetCPU()

	port, err := net.ResolveTCPAddr("tcp4", ":8000")
	base.CheckErr(err)
	listener, err := net.ListenTCP("tcp", port)
	base.CheckErr(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go waitTGW(conn)
	}
}
func waitTGW(conn net.Conn) {
	//tgw_l7_forward，处理
	buffer := make([]byte, TGW_HEADER_SIZE)
	length, err := conn.Read(buffer)
	if err != nil {
		log.Printf("read from client failed:%s", err.Error())
		return
	}
	segCount := 0
	var tgw []byte
	for i := 1; i < length; i++ {
		if buffer[i] == '\n' && buffer[i-1] == '\r' {
			segCount++
			if segCount == TGW_HEADER_SEG_COUNT {
				tgw = buffer[0 : i+1]
				buffer = buffer[i+1 : length]
				break
			}
		}
	}
	fmt.Println(string(tgw))
	fmt.Println("first pack after tgw (will pass):")
	base.TraceBytes2(buffer)
	go handleClient(conn)
}
func handleClient(conn net.Conn) {
	fmt.Println("A Client " + conn.RemoteAddr().String() + " in.")

	ch := make(chan []byte, 10)
	quit := make(chan int)

	go StartAgent(ch, conn, quit) //接收者

	header := make([]byte, 2)
	for {
		//header
		n, err := io.ReadFull(conn, header)
		if n == 0 && err == io.EOF {
			break
		} else if err != nil {
			log.Println("err read header:", err)
			goto OUT
		}

		//data
		size := binary.BigEndian.Uint16(header)
		body := make([]byte, size)
		//fmt.Println("\n\nheader info :", size)
		n, err = io.ReadFull(conn, body)
		if err != nil {
			log.Println("\nerr read body:", err)
			goto OUT
		}
		//base.TraceBytes2(body)

		ch <- body
	}
OUT:
	quit <- 0 //让agent来Close链接
}
