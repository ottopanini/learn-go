package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {

	var animal, action string

	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}

	for {
		fmt.Print(">")

		fmt.Scanln(&animal, &action)

		var animalInput Animal

		switch animal {
		case "cow":
			animalInput = cow
		case "snake":
			animalInput = snake
		case "bird":
			animalInput = bird
		}

		switch action {
		case "eat":
			animalInput.Eat()
		case "move":
			animalInput.Move()
		case "speak":
			animalInput.Speak()

		}

	}
}
