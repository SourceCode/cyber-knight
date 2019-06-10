package Narrator

import (
	"fmt"
	"github.com/SourceCode/cyber-knight/Player"
)



func MsgSelectClass() {
	fmt.Println("Select a class")
	for i, cName := range Player.ClassTypes {
		fmt.Printf("%d - %s\n", i + 1, cName)
	}
	fmt.Print("Enter 1-3: ")
}

func MsgEnterPlayerName() {
	fmt.Println("Please enter a player name")
}

func MsgEnterWorld(user *Player.UserClass) {
	println("Welcome to Cyber World young " + user.Class + ".  This world has never known a " + user.Name)
}

func MsgDescribeLocation(user *Player.UserClass) {
	println("You find yourself in a " + user.LocationName + " You look around and see: " + user.LocationDescription)
}

func MsgEnterAction(user *Player.UserClass) {
	println("Please enter an action such as forward, back, attack")
}

func MsgCommandList() {
	println("look - describes your environment")
	println("next - move to the next room")
	println("attack - attack the creature in the room")
}

func MsgCommandInvalid(a string) {
	println(a + " is not a valid action")
}