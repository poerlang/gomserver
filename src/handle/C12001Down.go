package handle

import (
	. "base"
)

type C12001Down struct {
	XX float32     //f32，横坐标
	ZZ float32     //f32，纵坐标
	YY float32     //f32，高度
	Dir float32    //f32，方向
	Action uint16  //u16，动作（静止、走路、奔跑、跑跳、原地跳、左横移、右横移、退后、退跑、攻击1、攻击2等等）
	Flag int8      //8，0不可移动，1成功移动
}

func (s *C12001Down)PackInTo(p *Pack) {
	p.WriteF32(s.XX)//横坐标
	p.WriteF32(s.ZZ)//纵坐标
	p.WriteF32(s.YY)//高度
	p.WriteF32(s.Dir)//方向
	p.WriteUInt16(s.Action)//动作（静止、走路、奔跑、跑跳、原地跳、左横移、右横移、退后、退跑、攻击1、攻击2等等）
	p.WriteInt8(s.Flag)//0不可移动，1成功移动
}
func (s *C12001Down)ToBytes() []byte {
	pack := NewPackEmpty()
	s.PackInTo(pack)
	return pack.Data()
}
