package Engine

type LootTable struct {
	Id int
	Biome string
	ItemId int
	ItemType string
	Act int
}


var LootTables = []LootTable{
	{1, "Room", 1, "Sword", 50},
	{1, "Room", 1, "Item", 1},
}