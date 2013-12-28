/**
 * Created by Administrator on 13-12-9.
 */
package main

import (
	"base"
	"net"
	"fmt"
)

func main() {

	defer base.Defer()

	base.SayHello("Gomo is runing.")
	base.SetCPU()

	port,err := net.ResolveTCPAddr("tcp4",":7981")
	base.CheckErr(err)
	listener,err := net.ListenTCP("tcp",port)
	base.CheckErr(err)

	for {
		conn,err := listener.Accept()
		if err!=nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn){
	fmt.Println(conn)
}
