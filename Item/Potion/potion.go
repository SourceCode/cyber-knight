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
	{Item.Item{1, "potion", "Health Potion", "Heals 50 health", 1 }, "+HP", "Heals 50 health", false, true},
}