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
			p := base.NewPack(data) //数据包裹
			c := p.ReadUInt16()     //协议号
			f := handle.DIC[c].Func //协议号对应函数
			fmt.Println("客户端请求协议：", c)
			if f != nil {
				b := f(c, p) //调用函数，得到结果
				if b != nil {
					sd <- b //发送
				}
			}
		case <-quit:
			close(quit)
			return
		}
	}
}
