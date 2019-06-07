package Player

type UserClass struct {
	Id int
	Class string
	Level int
	Name string
	Title string
	Hp int
	Mana int
	HpMax int
	ManaMax int
	Ap int
	ApMax int
	Str int
	Wis int
	Stm int
	LocationName string
	LocationDescription string
	LocationHash string
}

var ClassTypes = []string{"Techno Mage", "Data Lancer", "Byte Priest"}

func GetUserClass(id int) string {
	return ClassTypes[id - 1]
}