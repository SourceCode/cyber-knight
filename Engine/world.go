package Engine

import (
	"github.com/SourceCode/cyber-knight/Narrator"
	"github.com/SourceCode/cyber-knight/Player"
	"math/rand"
	"time"
)

type BiomeDescription struct {
	Name string
	Description string
}

var Biome = []string{"Room", "Court Yard", "Open Field"}

var BiomeDesc  = map[string][]BiomeDescription{
	Biome[0] : {
		{ "Dank Hallway", "The room is a hallway made of stench and rot."},
		{ "Royal Chamber", "The room is covered in royal decorations."},
		{"Torture Chamber", "The halls are covered bones in chains and blood stains."},
	},
	Biome[1] : {
		{ "Glass Rose Garden", "The courtyard is decorated in colorful roses made of glass."},
		{ "Training Ground", "The courtyard is used to train warriors."},
		{"Royal Garden", "The courtyard is planted with royal fruits and vegetables."},
	},
	Biome[2] : {
		{ "Makeshift Graveyard", "The field has been used as a makeshift graveyard ."},
		{ "Rotten Battleground", "The field is littered with corpses of the fallen."},
		{"Weeds and Dirt", "The field is littered in weeds and dirt piles."},
	},
}


func EnterGameWorld(user *Player.UserClass) {
	Narrator.MsgEnterWorld(user)
	EnterRandomGamePoint(user)
	GameLoop(user)
}

func EnterRandomGamePoint(user *Player.UserClass) {
	getRandomStartLocation(user)

}

func getRandomStartLocation(user *Player.UserClass) {
	rand.Seed(time.Now().Unix())
	id := rand.Intn(len(Biome))
	user.LocationName = Biome[id]
	bd := BiomeDesc[Biome[id]]
	roomId := rand.Intn(len(bd))
	user.LocationDescription = bd[roomId].Description
	Narrator.MsgDescribeLocation(user)
}