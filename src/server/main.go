/**
 * gomserver main.go
 */
package main

import (
	"base"
	"fmt"
	"net"
)

func main() {

	defer base.Defer()

	base.SayHello("gomserver is running.")
	base.SetCPU()

	port, err := net.ResolveTCPAddr("tcp4", ":7981")
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
	fmt.Println(conn)
}
