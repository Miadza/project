package animal

type Cat struct {
	ClimbTree string
}

func (c Cat) Sound() string {
	return "Мяу-мяу-мяу"
}

func (c Cat) Move() string {
	return "Кот может развивать скорость до 50 км/ч."
}

func (c Cat) Swim() int {
	return 1
}

func (c Cat) SoundFile() string {
	return "Cat.mp3"
}
