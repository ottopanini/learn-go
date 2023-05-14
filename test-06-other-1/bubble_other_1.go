package main

import (
	"fmt"
)

func Swap(numbers *[]int, j int) {
	(*numbers)[j], (*numbers)[j+1] = (*numbers)[j+1], (*numbers)[j]
}

func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				Swap(&numbers, j)
			}
		}
	}
	return numbers
}

func main() {
	var n []int
	for i := 0; i < 10; i++ {
		var j int
		_, err := fmt.Scan(&j)
		if err != nil {
			fmt.Println(err)
		}
		n = append(n, j)
	}
	fmt.Println(BubbleSort(n))
}
