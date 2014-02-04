package handle

import (
	. "base"
	"fmt"
)

type C20000Up struct {
	a1 []PosVO  //Array，包含[PosVO]
	a2 uint8    //u8，Just a val
}

func f20000Up(c uint16, p *Pack) []byte {
	s := new(C20000Up)
	count := int(p.ReadUInt16())             //数组长度（包含[PosVO]）
	for i := 0; i < count; i++ {
		node := new(PosVO)
		s.a1 = append(s.a1, node.UnPackFrom(p))
	}
	s.a2 = p.ReadUInt8()                     //Just a val

	fmt.Println(s)//需删除，否则影响性能
	return nil//需修改，返回不是nil的数据，确保客户端收到返回数据
}
