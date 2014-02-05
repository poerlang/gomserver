package handle

import (
	. "base"
	"fmt"
)

type C10000Up struct {
	SID string //String，
}

func f10000Up(c uint16, p *Pack) []byte {
	s := new(C10000Up)
	s.SID = p.ReadString()
	fmt.Println(s) //需删除，否则影响性能
	res := new(C10000Down)
	//向 res 赋值：
	res.Flag = 1
	return res.ToBytes()
}
