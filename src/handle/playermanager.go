package handle

import ()

const (
	MAX_PLAYERS = 20 //暂定最多20个玩家
)

var (
	AllPlayers map[string]*Player
)

func init() {
	AllPlayers = make(map[string]*Player)
}

func GetPlayer(key string) *Player {
	user, ok := AllPlayers[key]
	if ok {
		return user
	}
	return nil
}
func AddPlayer(key string, user *Player) *Player {
	user.State = 1
	user.SID = key
	AllPlayers[key] = user
	return user
}
func RemovePlayer(key string) {
	_, ok := AllPlayers[key]
	if ok {
		delete(AllPlayers, key)
	}
}
func (me *Player) SomeoneOffLine(u *Player) {
	if me == u {
		return
	}
	c := C12002Down{SID: u.SID}
	me.Sender <- c.ToBytes()
}
func (me *Player) SomeoneMove(u *Player) {
	if me == u {
		return
	}
	if u.State != 1 {
		return
	}
	c := C12001Down{SID: u.SID}
	c.XX = float32(u.XX)
	c.YY = float32(u.YY)
	c.ZZ = float32(u.ZZ)
	c.Dir = u.Dir
	c.Action = u.Action
	c.Flag = 1
	me.Sender <- c.ToBytes()
}
func (me *Player) SenderMsg(s string) {
	c := C11000Down{Str: s}
	me.Sender <- c.ToBytes()
}
