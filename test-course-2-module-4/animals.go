package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
}

func (c Cow) Eat() {
	println("grass")
}

func (c Cow) Move() {
	println("walk")
}

func (c Cow) Speak() {
	println("moo")
}

type Bird struct {
}

func (b Bird) Eat() {
	println("worms")
}

func (b Bird) Move() {
	println("fly")
}

func (b Bird) Speak() {
	println("peep")
}

type Snake struct {
}

func (s Snake) Eat() {
	println("mice")
}

func (s Snake) Move() {
	println("slither")
}

func (s Snake) Speak() {
	println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command, name, info string

		print("> ")
		fmt.Scan(&command, &name, &info)
		switch command {
		case "newanimal":
			created := false
			switch info {
			case "cow":
				animals[name] = Cow{}
				created = true
			case "bird":
				animals[name] = Bird{}
				created = true
			case "snake":
				animals[name] = Snake{}
				created = true
			default:
				println("Invalid animal name")
			}
			if created {
				println("Created it!")
			}
		case "query":
			switch info {
			case "eat":
				animals[name].Eat()
			case "move":
				animals[name].Move()
			case "speak":
				animals[name].Speak()
			default:
				println("Invalid animal info")
			}
		default:
			println("Invalid command")
		}
	}
}
