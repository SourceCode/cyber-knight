package Engine

import (
	"gitlab.com/rrentfro/go-sandbox-1/Narrator"
	"gitlab.com/rrentfro/go-sandbox-1/Player"
	"math/rand"
	"time"
)

type BiomeDesc struct {
	Name string
	Description string
}

var Biome = []string{"Dungeon", "Court Yard", "Open Field", "Garden", "Alley Way"}


func EnterGameWorld(user *Player.UserClass) {
	Narrator.MsgEnterWorld(user)
	EnterRandomGamePoint(user)
}

func EnterRandomGamePoint(user *Player.UserClass) {
	getRandomStartLocation(user)
	println("You find yourself in a " + user.LocationName)

	switch user.LocationName {

	}

}

func getRandomStartLocation(user *Player.UserClass) {
	rand.Seed(time.Now().Unix())
	user.LocationName = Biome[rand.Intn(len(Biome))]
}