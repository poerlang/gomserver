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

	//new一个Player但暂时不放入AllPlayers字典,等待用户发送10000后再放入字典，
	//字典的 key为10000传送过来的SID。
	user := new(handle.Player)
	user.SID = ""
	user.Conn = conn
	user.State = 0

	go StartSender(sd, conn) //发送者
	for {
		select {
		case data := <-c:
			p := base.NewPack(data) //读取数据包裹
			c := p.ReadUInt16()     //读取协议号
			f := handle.DIC[c].Func //获得协议号对应函数
			fmt.Println("客户端请求协议：", c)
			if f != nil {
				b := f(c, p, user) //调用函数，得到结果
				if b != nil {
					pp := base.NewPackEmpty()
					pp.WriteUInt16(c)   //写入协议号
					pp.WriteRawBytes(b) //写入返回数据
					sd <- pp.Data()     //发送
				}
			}
		case <-quit:
			close(quit)
			//todo:发送断开链接的警告
			//todo:其他清理工作
			user.State = -1
			user.Conn.Close()
			handle.RemovePlayer(user.SID)
			return
		}
	}
}
