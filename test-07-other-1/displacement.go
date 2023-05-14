package main

import "fmt"

func main() {
	parameters := [4]float64{}
	fmt.Println("Enter acceleration, initial velocity, initial displacement and time(seconds), separated by space: ")
	for i := 0; i < 4; i++ {
		fmt.Scan(&parameters[i])
	}
	fn := GenDisplaceFn(parameters[0], parameters[1], parameters[2])
	fmt.Print("Displacement after time: ", fn(parameters[3]))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*t*t + v*t + s
	}
}
