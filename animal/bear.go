package animal

type Bear struct {
	HighSpeed string
}

func (b Bear) Sound() string {
	return "Медведь может рычать, свистеть, пищать, фыркать."
}

func (b Bear) Move() string {
	return "Медведь является очень быстрым"
}

func (b Bear) Swim() int {
	return 1
}
