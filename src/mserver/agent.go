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
	sd_quit := make(chan int)

	//new一个Player但暂时不放入AllPlayers字典,等待用户发送10000后再放入字典，
	//字典的 key为10000传送过来的SID。
	user := new(handle.Player)
	user.SID = ""
	user.Conn = conn
	user.State = 0
	user.Sender = sd

	go StartSender(sd, conn, sd_quit, &user) //发送者
	for {
		select {
		case data := <-c:
			p := base.NewPack(data) //读取数据包裹
			c := p.ReadUInt16()     //读取协议号
			f := handle.DIC[c].Func //获得协议号对应函数
			//fmt.Println("客户端请求协议：", c)
			if f != nil {
				b := f(c, p, user) //调用函数，得到结果
				if b != nil {
					sd <- b //发送
				}
			}
		case <-quit:
			close(quit)
			fmt.Println("断开用户：" + user.SID)

			//告诉周围其他玩家，此user下线了。
			near := handle.MapA.Tree.FindNearObjects_RLq(user.GetPreviousPos(), 100)
			for _, o := range near {
				other, ok := o.(*handle.Player)
				fmt.Println(other.SID + " " + other.Map)
				if !ok {
					continue
				}
				other.SomeoneOffLine(user)
			}
			//从地图中移除此玩家
			if user.Map == "MapA" {
				handle.MapA.Tree.Remove_WLq(user)
			}

			//todo:其他清理工作

			user.State = -1
			user.Conn.Close()
			fmt.Println("用户【" + user.SID + "】已断开")
			user.Conn = nil
			handle.RemovePlayer(user.SID)
			user = nil
			return
		}
	}
}
