package handle

import (
	. "base"
)

type C11000Down struct {
	Str string //String，消息内容
}

func (s *C11000Down) PackInTo(p *Pack) {
	p.WriteUInt16(11000) //写入协议号
	p.WriteString(s.Str) //消息内容
}
func (s *C11000Down) ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
