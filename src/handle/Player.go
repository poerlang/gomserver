package handle

import (
	. "base"
)

type Player struct {
	SID      string //String，随机字符串
	NickName string //String，用户自定义名
	State    int8   //8，状态，-1下线，0上线但尚未通过登录验证、或未初始化，1在线
	Map      string //String，所在地图
}

func (s *Player) UnPackFrom(p *Pack) Player {
	s.SID = p.ReadString()      //随机字符串
	s.NickName = p.ReadString() //用户自定义名
	s.State = p.ReadInt8()      //状态，-1下线，0上线但尚未通过登录验证、或未初始化，1在线
	s.Map = p.ReadString()      //所在地图
	return *s
}

func (s *Player) PackInTo(p *Pack) {
	p.WriteString(s.SID)      //随机字符串
	p.WriteString(s.NickName) //用户自定义名
	p.WriteInt8(s.State)      //状态，-1下线，0上线但尚未通过登录验证、或未初始化，1在线
	p.WriteString(s.Map)      //所在地图
}

func (s *Player) ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
