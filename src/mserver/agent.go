// agent
package main

import (
	"base"
	"fmt"
	"handle"
	"net"
)

func StartAgent(c chan []byte, conn net.Conn, quit chan int) {
	sd := make(chan []byte, 10)
	go StartSender(sd, conn) //发送者
	for {
		select {
		case data := <-c:
			p := base.NewPack(data) //读取数据包裹
			c := p.ReadUInt16()     //读取协议号
			f := handle.DIC[c].Func //获得协议号对应函数
			fmt.Println("客户端请求协议：", c)
			if f != nil {
				b := f(c, p) //调用函数，得到结果
				if b != nil {
					pp := base.NewPackEmpty()
					pp.WriteUInt16(c)   //写入协议号
					pp.WriteRawBytes(b) //写入返回数据
					sd <- pp.Data()     //发送
				}
			}
		case <-quit:
			close(quit)
			return
		}
	}
}
