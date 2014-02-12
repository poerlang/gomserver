package handle

import (
	. "base"
	//"fmt"
	. "twof"
)

type C12001Up struct {
	XX     float32 //f32，横坐标
	ZZ     float32 //f32，纵坐标
	YY     float32 //f32，高度
	Dir    float32 //f32，方向
	Action uint16  //u16，动作（静止、走路、奔跑、跑跳、原地跳、左横移、右横移、退后、退跑、攻击1、攻击2等等）
}

func f12001Up(c uint16, p *Pack, u *Player) []byte {
	s := new(C12001Up)
	s.XX = p.ReadF32()        //横坐标
	s.ZZ = p.ReadF32()        //纵坐标
	s.YY = p.ReadF32()        //高度
	s.Dir = p.ReadF32()       //方向
	s.Action = p.ReadUInt16() //动作（静止、走路、奔跑、跑跳、原地跳、左横移、右横移、退后、退跑、攻击1、攻击2等等）
	//业务逻辑：
	if u.State != 1 {
		return nil
	}
	u.XX = float64(s.XX)
	u.YY = float64(s.YY)
	u.ZZ = float64(s.ZZ)
	u.Dir = s.Dir
	u.Action = s.Action

	xz := &TwoF{float64(u.XX), float64(u.ZZ)}
	MapA.Tree.Move_WLq(u, xz)

	//通知周围其他玩家
	nears := MapA.Tree.FindNearObjects_RLq(xz, 1000)
	for _, o := range nears {
		other, ok := o.(*Player)
		if !ok {
			continue
		}
		other.SomeoneMove(u)
	}
	return nil //不需要回传消息给玩家自己
}
