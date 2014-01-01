// agent
package main

import (
	"base"
	"fmt"
	"net"
)

func StartAgent(c chan []byte, conn net.Conn) {
	for {
		data, ok := <-c
		if ok {
			p := base.NewPack(data)
			fmt.Println(p.ReadInt16()) //cmd
			fmt.Println("a:  ", p.ReadDouble642())
			fmt.Println("b:  ", p.ReadString())
		}
	}
}
