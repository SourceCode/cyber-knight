package Engine

import (
	"github.com/SourceCode/cyber-knight/Player"
)

type LocationMap struct {
	Id int
	LocationName string
	LocationDescription string
}

type GameState struct {
	Map map[string][]LocationMap
}

type SysState struct {
	State []LocationMap
}

var LiveState SysState

func PushState (user *Player.UserClass) {
	s := LocationMap{len(LiveState.State), user.LocationName, user.LocationDescription}
	LiveState.State = append(LiveState.State, s)
}