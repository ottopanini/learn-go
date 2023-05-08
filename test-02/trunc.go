package main

import "fmt"

// rename to main
func main2() {
	var number float64

	fmt.Printf("Input a string: ")
	fmt.Scan(&number)
	fmt.Printf("Truncated: %d\n", int(number))
}
