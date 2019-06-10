package Engine

import (
	"github.com/SourceCode/cyber-knight/Narrator"
	"os"
)

func RouteAction(action string) bool {

	switch action {
	case "help":
		Narrator.MsgCommandList()
		return true
	case "look":

	case "bag":

	case "pickup":

	case "attack":

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