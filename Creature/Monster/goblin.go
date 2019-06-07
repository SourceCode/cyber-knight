package Monster

type Goblin struct {
	currentHp int
}

func (g *Goblin) MaxHp() int {
	return 50
}

func (g *Goblin) Name() string {

	return "Goblin"
}

func (g *Goblin) Evil() bool {

	return true
}

func (g *Goblin) Level() int {

	return 1
}

func (g *Goblin) Hp() int {

	return 50
}