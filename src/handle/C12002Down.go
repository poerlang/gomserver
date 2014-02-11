package handle

import (
	. "base"
)

type C12002Down struct {
	SID string //String，掉线玩家的SID
}

func (s *C12002Down) PackInTo(p *Pack) {
	p.WriteUInt16(12002) //写入协议号
	p.WriteString(s.SID) //掉线玩家的SID
}
func (s *C12002Down) ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
