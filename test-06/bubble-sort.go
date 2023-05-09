package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter integers (each submitted by ENTER, Empty values will be handled as end of input): ")

	//r, _ := io.Pip1e()
	a := input(bufio.NewReader(os.Stdin))
	BubbleSort(a)
	fmt.Println(a)
}

func input(in *bufio.Reader) []int {
	scanner := bufio.NewScanner(in)
	var n int
	a := make([]int, n)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}
		var err error
		number, err := strconv.Atoi(line)

		if err != nil {
			fmt.Println(err)
			continue
		}
		a = append(a, number)
	}
	return a
}

func BubbleSort(a []int) {
	swapped := true
	for n := len(a); swapped; n-- {
		swapped = false
		for i := 1; i < n; i++ {
			if a[i-1] > a[i] {
				Swap(a, i-1)
				swapped = true
			}
		}
	}
}

func Swap(a []int, i int) {
	a[i], a[i+1] = a[i+1], a[i]
}
