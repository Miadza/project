package main

import (
	"fmt"

	"helloapp/animal"
	"helloapp/interaction"
	"helloapp/sound"

	"github.com/gen2brain/beeep"
)

func main() {


	var inputType string
	fmt.Println("Введите тип животных, о каких хотите узнать (пример: хищные, птицы, млекопитающие)")
	fmt.Scanln(&inputType)
	if _, err := interaction.Errk(inputType); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Все топ")
	}

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
			sound.PlaySound(bear)
			err := beeep.Notify("Title", "Message body", "assets/information.png")
			if err != nil {
				panic(err)
			}
			// 	case 2:
			// 		smallPanda := animal.SmallPanda{Size: "Длина тела 51–64 см."}
			// 		fmt.Println("Информация о малой панде:")
			// 		interaction.PrintAnimal(smallPanda)
			// 		fmt.Println("Уникальная информация:", smallPanda.Size)
			// 	case 3:
			// 		cat := animal.Cat{ClimbTree: "Умеет лазать по деревьям"}
			// 		fmt.Println("Информация о коте:")
			// 		interaction.PrintAnimal(cat)
			// 		fmt.Println("Уникальная информация:", cat.ClimbTree)
			// 		interaction.PrintSwim(cat)
			// 	case 4:
			// 		dog := animal.Dog{RecognizeDiseases: "Распознает болезни"}
			// 		fmt.Println("Информация о собаке:")
			// 		interaction.PrintAnimal(dog)
			// 		fmt.Println("Уникальная информация:", dog.RecognizeDiseases)
			// 		interaction.PrintSwim(dog)
			// 	}
			// }
			// switch inputType {
			// case "птицы":
			// 	fmt.Println("О ком вы хотите узнать (введите цифру)")
			// 	fmt.Println("1. Пингвин")
			// 	var input int
			// 	fmt.Scanf("%d", &input)
			// 	switch input {
			// 	case 1:
			// 		penguin := animal.Penguin{CantFly: "Единственная птица не умеет летать"}
			// 		fmt.Println("Информация о пингвине:")
			// 		interaction.PrintAnimal(penguin)
			// 		fmt.Println("Уникальная информация:", penguin.CantFly)
			// 		interaction.PrintSwim(penguin)
			// 	}
			// }
			// switch inputType {
			// case "млекопитающие":
			// 	fmt.Println("О ком вы хотите узнать (введите цифру)")
			// 	fmt.Println("1. Жираф")
			// 	var input int
			// 	fmt.Scanf("%d", &input)
			// 	switch input {
			// 	case 1:
			// 		giraffe := animal.Giraffe{Noswimming: "Не умеет плавать"}
			// 		fmt.Println("Информация о жирафе:")
			// 		interaction.PrintAnimal(giraffe)
			// 		fmt.Println("Уникальная информация:", giraffe.Noswimming)
			// 		interaction.PrintSwim(giraffe)
			// 	}
			// }
		}
	}
}
