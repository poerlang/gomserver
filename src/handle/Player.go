package handle

import (
	"net"
	"twof"
)

type Player struct {
	SID      string   //String，随机字符串
	ID       uint32   //序号
	NickName string   //String，用户自定义名
	State    int8     //8，状态，-1下线，0上线但尚未通过登录验证、或未初始化，1在线
	Map      string   //String，所在地图
	XX       float64  //横坐标
	ZZ       float64  //纵坐标
	YY       float64  //高度
	Dir      float32  //方向
	Action   uint16   //动作
	Speed    float32  //速度
	Conn     net.Conn //链接
}

func (u *Player) GetPreviousPos() *twof.TwoF { return &twof.TwoF{u.XX, u.ZZ} } // Always 0 for testing
func (u *Player) GetId() uint32              { return u.ID }                   // Not needed for testing
func (u *Player) GetType() uint8             { return 0 }                      // Not needed for testing
func (u *Player) GetZ() float64              { return u.YY }                   // Not needed for testing
func (u *Player) GetDir() float32            { return u.Dir }                  // Not needed for testing
