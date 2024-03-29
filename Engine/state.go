package Engine

import (
	"github.com/SourceCode/cyber-knight/Player"
)

type LocationMap struct {
	Id          int
	Name        string
	Description string
}

type PlayerItem struct {
	Id int
	Name string
	Description string
	Act int
}

type SysState struct {
	State []LocationMap
	Bag map[string][]PlayerItem
	Gold int
}

var LiveState = SysState{Bag: map[string][]PlayerItem{
	"Items": {{0, "Bomb", "A fun explosive", 3}} }, Gold: 0}

func PushState (user *Player.UserClass) {
	s := LocationMap{len(LiveState.State), user.LocationName, user.LocationDescription}
	LiveState.State = append(LiveState.State, s)
}