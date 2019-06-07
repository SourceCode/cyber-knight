package main

import (
	"bufio"
	"fmt"
	"github.com/SourceCode/cyber-knight/Engine"
	"github.com/SourceCode/cyber-knight/Narrator"
	"github.com/SourceCode/cyber-knight/Player"
	"os"
	"strconv"
	"strings"
)

var user = Player.UserClass{Id: 1, Level: 1, Hp: 100, Mana: 100, HpMax: 100, ManaMax: 100, Ap: 10, ApMax: 20, Str: 1, Wis: 1, Stm: 5}

func main() {
	fmt.Println("Cyber Knight - 0.1.0")

	reader := bufio.NewReader(os.Stdin)

	Narrator.MsgSelectClass()

	cs, _ := reader.ReadString('\n')
	cs = strings.TrimSuffix(cs, "\n")
	csId, _ := strconv.Atoi(cs)

	// We should verify input
	user.Class = Player.GetUserClass(csId)
	user.Id = csId

	//Start Game
	//fmt.Println(csId, error)
	//fmt.Println(user)

	Engine.EnterGameWorld(&user)
}
