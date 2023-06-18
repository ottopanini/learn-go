package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Example commands:")
	fmt.Println("> newanimal {name} {cow|bird|snake}")
	fmt.Println("> query {name} {eat|move|speak}")
	fmt.Println()
	animals := make(map[string]Animal)
	for {
		fmt.Printf("> ")
		s := readConsoleLine()
		w := strings.Fields(s)
		if len(w) != 3 {
			fmt.Println("Wrong command, expected 3 arguments")
			continue
		}
		command := w[0]
		switch command {
		case "newanimal":
			name := w[1]
			animalType := w[2]
			var animal Animal
			switch animalType {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Unsupported animal type")
				continue
			}
			animals[name] = animal
			fmt.Println("Created it!")
		case "query":
			name := w[1]
			action := w[2]
			animal, ok := animals[name]
			if !ok {
				fmt.Println("Animal name not found")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Unsupported action")
				continue
			}
		default:
			fmt.Println("Unsupported command")
			continue
		}
	}
}

func readConsoleLine() string {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	s = strings.Replace(s, "\n", "", 1)
	return s
}

// Animal interface abstracts the methods Eat, Move and Speak
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow implements interface Animal
type Cow struct{}

// Eat prints the cow's food
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move prints the cow's locomotion
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak prints the cow's sound
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird implements interface Animal
type Bird struct{}

// Eat prints the bird's food
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move prints the bird's locomotion
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak prints the bird's sound
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake implements interface Animal
type Snake struct{}

// Eat prints the snake's food
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move prints the bird's locomotion
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak prints the snake's sound
func (s Snake) Speak() {
	fmt.Println("hsss")
}
