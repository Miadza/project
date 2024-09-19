package animal

type Penguin struct {
	CantFly string
}

func (p Penguin) Sound() string {
	return "На суше они общаются криками, напоминающими звуки трубы."
}

func (p Penguin) Move() string {
	return "Средняя скорость в воде — 5–10 км/ч."
}

func (p Penguin) Swim() int {
	return 1
}
