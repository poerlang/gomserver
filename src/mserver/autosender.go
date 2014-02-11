// agent
package main

import (
	//"base"
	//"fmt"
	"handle"
	//"strconv"
	"time"
)

func StartAutoSender(sd chan []byte, asd_quit chan int, u *handle.Player) {
	tiker := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		case <-tiker.C:
			nears := handle.MapA.Tree.FindNearObjects_RLq(u.GetPreviousPos(), 1000)
			for _, o := range nears {
				other, ok := o.(*handle.Player)
				if !ok {
					continue
				}
				other.SomeoneMove(u)
			}
		case <-asd_quit:
			close(asd_quit)
			return
		}
	}
}
