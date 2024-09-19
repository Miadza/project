package main

import (
	"fmt"
	"helloapp/animal"
	"helloapp/interaction"

	"rsc.io/quote"
)

func main() {
	message := quote.Hello()
	fmt.Println(message)

	var inputType string
	fmt.Println("Введите тип животных, о каких хотите узнать (пример: хищные, птицы, млекопитающие)")
	fmt.Scanln(&inputType)

	switch inputType {
	case "хищные":
		fmt.Println("1. Медведь")
		fmt.Println("2. Малая панда")
		fmt.Println("3. Кот")
		fmt.Println("4. Собака")
		fmt.Println("О ком вы хотите узнать (введите цифру)")
		var input int
		fmt.Scanf("%d", &input)

		switch input {
		case 1:
			bear := animal.Bear{HighSpeed: "Достигает 60 км/ч"}
			fmt.Println("Информация о медведе:")
			interaction.PrintAnimal(bear)
			fmt.Println("Уникальная информация:", bear.HighSpeed)
			interaction.PrintSwim(bear)
		case 2:
			smallPanda := animal.SmallPanda{Size: "Длина тела 51–64 см."}
			fmt.Println("Информация о малой панде:")
			interaction.PrintAnimal(smallPanda)
			fmt.Println("Уникальная информация:", smallPanda.Size)
		case 3:
			cat := animal.Cat{ClimbTree: "Умеет лазать по деревьям"}
			fmt.Println("Информация о коте:")
			interaction.PrintAnimal(cat)
			fmt.Println("Уникальная информация:", cat.ClimbTree)
			interaction.PrintSwim(cat)
		case 4:
			dog := animal.Dog{RecognizeDiseases: "Распознает болезни"}
			fmt.Println("Информация о собаке:")
			interaction.PrintAnimal(dog)
			fmt.Println("Уникальная информация:", dog.RecognizeDiseases)
			interaction.PrintSwim(dog)
		}
	}
}
