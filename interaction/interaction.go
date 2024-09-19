package interaction

import (
	"fmt"
	"helloapp/animal"
)

func PrintAnimal(a animal.Animal) {
	fmt.Println("Звуки:", a.Sound())
	fmt.Println("Движение:", a.Move())
}

func PrintSwim(s animal.Swimmer) {
	var canSwim = s.Swim()
	if canSwim == 1 {
		fmt.Println("Плавает")
	} else {
		fmt.Println("Не умеет плавать")
	}
}
