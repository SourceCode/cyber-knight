package Engine

import (
	"github.com/SourceCode/cyber-knight/Narrator"
	"github.com/SourceCode/cyber-knight/Player"
	"os"
)

func RouteAction(action string, user *Player.UserClass) bool {

	switch action {
	case "help":
		Narrator.MsgCommandList()
		return true
	case "look":
		Narrator.MsgDescribeLocation(user)
		return true
	case "bag":

	case "pickup":

	case "attack":

	case "forward":
		Narrator.MsgMoveNextRoom()
		MoveNextRoom(user)
		return true
	case "back":

	case "quit":
		os.Exit(3)
		return true
	case "exit":
		os.Exit(3)
		return true
	}

	Narrator.MsgCommandInvalid(action)
	return true
}