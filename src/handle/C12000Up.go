package handle

import (
	. "base"
	//"fmt"
	"twof"
)

type C12000Up struct {
	MapName string //String，地图名字
}

func f12000Up(c uint16, p *Pack, u *Player) []byte {
	s := new(C12000Up)
	s.MapName = p.ReadString() //地图名字
	//fmt.Println(s)             //需删除，否则影响性能
	res := new(C12000Down)
	//业务逻辑：
	pos := twof.TwoF{0, 0}
	if u.Map == "" {
		MapA.Tree.Add_WLq(u, &pos)
		u.Map = s.MapName
		res.Flag = 1
	}
	return res.ToBytes()
}
