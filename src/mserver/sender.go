// agent
package main

import (
	"base"
	"fmt"
	"net"
)

func StartSender(c chan []byte, conn net.Conn) {
	for {
		data, ok := <-c
		if ok {
			l := len(data) //长度
			fmt.Println("\t即将发给前端 ", l, " 字节")
			p := base.NewPackEmpty()       //空数据包裹
			p.WriteUInt16(uint16(l))       //写入长度
			p.WriteRawBytes(data)          //写入数据
			p.TraceBytes()                 //打印
			_, err := conn.Write(p.Data()) //输出到链接
			if err != nil {
				fmt.Println("Error send reply :", err)
				return
			}
		}
	}
}
