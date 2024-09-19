package animal

type SmallPanda struct {
	Size string
}

func (s SmallPanda) Sound() string {
	return "Малые панды издают звуки, напоминающие птичье щебетание."
}

func (s SmallPanda) Move() string {
	return "Малые панды ловко лазают по веткам."
}

func (s SmallPanda) Swim() int {
	return 1
}
