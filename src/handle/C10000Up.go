package handle

import (
	. "base"
)

type C10000Up struct {
	SID string //String，
}

func f10000Up(c uint16, p *Pack, u *Player) []byte {
	s := new(C10000Up)
	s.SID = p.ReadString()
	res := new(C10000Down)
	//业务逻辑：
	AddPlayer(s.SID, u)
	res.Flag = 1
	return res.ToBytes()
}
