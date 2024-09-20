package animal

type Bear struct {
	HighSpeed string
}

func (b Bear) Sound() string {
	return "ррр"
}

func (b Bear) Move() string {
	return "Медведь является очень быстрым"
}

func (b Bear) Swim() int {
	return 1
}

func (b Bear) SoundFile() string {
	return "Bear.mp3"
}
