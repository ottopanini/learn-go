package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result := GenDisplaceFn(promptInitialValues(reader))(promptTime(reader))
	fmt.Printf("Displacement: %v\n", result)
}

func GenDisplaceFn(acceleration float64, initVelocity float64, initDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initVelocity*time + initDisplacement
	}
}

func promptInitialValues(in *bufio.Reader) (float64, float64, float64) {
	fmt.Println("Enter acceleration, initial velocity and initial displacement (each submitted by ENTER): ")
	return promptFloat64(in), promptFloat64(in), promptFloat64(in)
}

func promptTime(in *bufio.Reader) float64 {
	fmt.Println("Enter time: ")
	return promptFloat64(in)
}

func promptFloat64(in *bufio.Reader) float64 {
	var err = errors.New("initial")
	var float float64
	for err != nil {
		var line []byte
		line, _, err = in.ReadLine()
		if err != nil {
			fmt.Println(err)
			continue
		}
		float, err = strconv.ParseFloat(string(line), 64)
		if err != nil {
			fmt.Println(err)
		}
	}
	return float
}
