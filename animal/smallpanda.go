package animal

type Smallpanda struct {
	Size string
}

func (s Smallpanda) Sound() string {
	return "Малые панды издают звуки, напоминающие птичье щебетание."
}

func (s Smallpanda) Move() string {
	return "Малые панды ловко лазают по веткам."
}

func (s Smallpanda) Swim() int {
	return 1
}

func (s Smallpanda) SoundFile() string {
	return "Panda.mp3"
}