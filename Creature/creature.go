package Creature

type Creature interface {
	Name() string
	Evil() bool
	Level() int
	Hp() int
	MaxHp() int
}

