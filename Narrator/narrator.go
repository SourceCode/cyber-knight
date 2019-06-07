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

func MsgEnterWorld(user *Player.UserClass) {
	println("Welcome to Cyber World young " + user.Class)
}

func MsgDescribeLocation(user *Player.UserClass) {
	println("You find yourself in a " + user.LocationName)


}