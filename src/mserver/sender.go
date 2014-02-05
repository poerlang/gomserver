// agent
package main

import (
	"base"
	"fmt"
	"net"
)

func StartSender(c chan []byte, conn net.Conn) {
	for {
	OUT:
		data, ok := <-c
		if ok {
			l := len(data) //长度
			fmt.Println("\n\t即将发给前端 ", l, " 字节")
			p := base.NewPackEmpty() //空数据包裹
			p.WriteUInt16(uint16(l)) //写入长度
			p.WriteRawBytes(data)    //写入数据
			p.TraceBytes()           //打印
			b := p.Data()            //[]byte
			for len(b) > 0 {
				n, err := conn.Write(b) //输出到链接
				if err == nil {
					if n != len(b) {
						fmt.Println("没有一次性写入")
					}
					b = b[n:]
				} else if e2, ok := err.(*net.OpError); ok && (e2.Temporary() || e2.Timeout()) {
					continue
				} else {
					fmt.Println("网络问题")
					goto OUT
				}
			}
		}
	}
}
