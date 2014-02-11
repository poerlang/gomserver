/**
 * gomserver main.go
 */
package main

import (
	"base"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

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
		go handleClient(conn)
	}
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
		//fmt.Println("\n\n收到客户端消息，bodySize:", size)
		n, err = io.ReadFull(conn, body)
		if err != nil {
			log.Println("\nerr read body:", err)
			goto OUT
		}
		ch <- body
	}
OUT:
	quit <- 0 //让agent来Close链接
}
