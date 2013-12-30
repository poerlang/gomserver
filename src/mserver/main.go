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
	defer conn.Close()
	fmt.Println(conn.RemoteAddr().String() + " in.")
	header := make([]byte, 2)
	ch := make(chan []byte, 10)
	go StartAgent(ch, conn)
	for {
		//header
		n, err := io.ReadFull(conn, header)
		if n == 0 && err == io.EOF {
			break
		} else if err != nil {
			log.Println("err read header:", err)
			break
		}

		//data
		size := binary.BigEndian.Uint16(header)
		body := make([]byte, size)
		fmt.Println("header:", n, " size:", size)
		n, err = io.ReadFull(conn, body)
		if err != nil {
			log.Println("err read body:", err)
			break
		}
		ch <- body
	}
	close(ch)
}
