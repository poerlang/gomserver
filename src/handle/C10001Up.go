package handle

import (
	. "base"
	"fmt"
)

type C10001Up struct {
	SID string //String，玩家唯一标识
}

func f10001Up(c uint16, p *Pack, u *Player) []byte {
	s := new(C10001Up)
	s.SID = p.ReadString() //玩家唯一标识
	fmt.Println(s)         //需删除，否则影响性能
	res := new(C10001Down)
	//业务逻辑：
	user := GetPlayer(s.SID)
	if user != nil {
		res.Flag = 1
		fmt.Println("玩家存在")
	} else {
		fmt.Println("玩家不存在")
	}
	return res.ToBytes()
}
