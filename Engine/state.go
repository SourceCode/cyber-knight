package Engine

type LocationMap struct {
	X int
	Y int
	LocationName string
	LocationDescription string
}

type GameState struct {
	Map map[string][]LocationMap
	
}

