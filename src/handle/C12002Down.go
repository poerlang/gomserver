package handle

import (
	. "base"
)

type C12002Down struct {
	SID string  //String，掉线玩家的SID
}

func (s *C12002Down)PackInTo(p *Pack) {
	p.WriteString(s.SID)//掉线玩家的SID
}
func (s *C12002Down)ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
