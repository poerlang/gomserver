package handle

import (
	. "base"
)

type PosVO struct {
	arr []uint8  //Array，包含[u8]
	str string   //String，Just a str
}

func (s *PosVO) UnPackFrom(p *Pack) PosVO {
	count := int(p.ReadUInt16())          //数组长度（包含[u8]）
	for i := 0; i < count; i++ {
		s.arr = append(s.arr, p.ReadUInt8())
	}
	s.str = p.ReadString()                //Just a str
	return *s
}
func (s *PosVO) PackInTo(p *Pack) {
	count := len(s.arr)          //数组长度（包含[u8]）
	for i := 0; i < count; i++ {
		p.WriteUInt8(s.arr[i])
	}
	p.WriteString(s.str)         //Just a str
}
