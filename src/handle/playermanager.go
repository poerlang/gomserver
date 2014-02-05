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
