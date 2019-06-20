package Item

import "github.com/SourceCode/cyber-knight/Item"

type Potion struct {
	Item Item.Item
	Effect string
	EffectDesc string
	Unique bool
	Stack bool
}


var PotionTable = []Potion{
	{{1, "potion", "Health Potion", }}
}