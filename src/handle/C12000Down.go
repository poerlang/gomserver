package handle

import (
	. "base"
)

type C12000Down struct {
	MapName string //String，地图名字
	Flag    int8   //8，1进入成功，0地图不存在
}

func (s *C12000Down) PackInTo(p *Pack) {
	p.WriteUInt16(12000)     //写入协议号
	p.WriteString(s.MapName) //地图名字
	p.WriteInt8(s.Flag)      //1进入成功，0地图不存在
}
func (s *C12000Down) ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
