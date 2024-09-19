package animal

type Dog struct {
	RecognizeDiseases string
}

func (d Dog) Move() string {
	return "Обычно собаки могут бегать со скоростью 10–15 км/ч."
}

func (d Dog) Sound() string {
	return "Гав-гав"
}

func (d Dog) Swim() int {
	return 1
}
