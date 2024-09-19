package animal

type Swimmer interface {
	Swim() int
}

type Animal interface {
	Sound() string
	Move() string
}
