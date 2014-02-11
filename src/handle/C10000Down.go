package handle

import (
	. "base"
)

type C10000Down struct {
	Flag int8 //8，1登录成功，0登录失败
}

func (s *C10000Down) PackInTo(p *Pack) {
	p.WriteUInt16(10000) //写入协议号
	p.WriteInt8(s.Flag)  //1登录成功，0登录失败
}
func (s *C10000Down) ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
