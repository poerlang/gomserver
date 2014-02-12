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
			//todo:.......
		case <-asd_quit:
			close(asd_quit)
			return
		}
	}
}
