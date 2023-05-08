package main

import (
	"fmt"
	"strconv"
)

func main2() {
	numbers := make([]int, 0, 3)
	for input := ""; input != "X"; {
		fmt.Printf("Input an integer (Type 'X' to leave): \n")
		fmt.Scan(&input)

		if input == "X" {
			break
		}

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		insertOffset := 0
		insertedNumbers := make([]int, len(numbers)+1)
		if len(numbers) == 0 {
			insertedNumbers[0] = number
		} else {
			if number < numbers[0] {
				insertedNumbers[0] = number
				insertOffset = 1
			} else if number > numbers[len(numbers)-1] {
				insertedNumbers[len(numbers)] = number
			}

			for i, n := range numbers {
				if number < n && insertOffset == 0 {
					insertOffset = 1
					insertedNumbers[i] = number
				}

				insertedNumbers[i+insertOffset] = numbers[i]
			}
		}
		numbers = insertedNumbers

		fmt.Printf("Sorted Numbers: ")
		for i, n := range numbers {
			fmt.Printf("%d", n)

			if i < len(numbers)-1 {
				fmt.Printf(", ")
			} else {
				fmt.Printf("\n")
			}
		}
	}
}
