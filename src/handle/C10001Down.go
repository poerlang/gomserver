package handle

import (
	. "base"
)

type C10001Down struct {
	Flag int8 //8，0不在线，1在线
}

func (s *C10001Down) PackInTo(p *Pack) {
	p.WriteUInt16(10001) //写入协议号
	p.WriteInt8(s.Flag)  //0不在线，1在线
}
func (s *C10001Down) ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
