package main

import "fmt"

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	for {
		var animal, info string
		fmt.Print("> ")
		fmt.Scan(&animal, &info)
		fmt.Println(getValue(animals[animal], info))
	}
}

func getValue(animal Animal, property string) string {
	switch property {
	case "eat":
		return animal.food
	case "move":
		return animal.locomotion
	case "speak":
		return animal.noise
	}
	return ""
}
