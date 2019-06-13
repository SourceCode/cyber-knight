package Engine

import (
	"bufio"
	"fmt"
	"github.com/SourceCode/cyber-knight/Narrator"
	"github.com/SourceCode/cyber-knight/Player"
	"os"
	"strings"
)

func GameLoop(user *Player.UserClass) {

	for {
		Narrator.MsgEnterAction(user)
		reader := bufio.NewReader(os.Stdin)
		action, _ := reader.ReadString('\n')
		action = strings.TrimSuffix(action, "\n")
		RouteAction(action, user)
		fmt.Printf("The action was %s\n", action)

	}


}
