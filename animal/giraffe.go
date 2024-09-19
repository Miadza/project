package animal

type Giraffe struct {
	Noswimming string
}

func (g Giraffe) Sound() string {
	return "Хотя жирафы обычно молчаливы, они могут издавать разные звуки."
}

func (g Giraffe) Move() string {
	return "Скорость бега жирафа может достигать 60 км/ч"
}

func (g Giraffe) Swim() int {
	return 0
}
