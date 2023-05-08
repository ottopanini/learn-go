package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main2() {
	fmt.Println("enter text:")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("User input is: ", input)

	input = strings.ToLower(input)

	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Not found!")
	}

	if strings.ContainsRune(input, 'i') && strings.ContainsRune(input, 'a') && strings.ContainsRune(input, 'n') {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
