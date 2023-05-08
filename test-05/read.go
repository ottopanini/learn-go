package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {
	file := getFile()
	fileContent := readFile(file)
	names := readNames(fileContent)
	printNames(names)
}

func getFile() string {
	var file string
	fmt.Print("Enter file name: ")
	_, err := fmt.Scanf("%s", &file)
	if err != nil {
		panic(err)
	}
	return file
}

func readFile(file string) string {
	contents, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(contents)
}

func readNames(fileContent string) []Person {
	var names []Person
	scanner := bufio.NewScanner(strings.NewReader(fileContent))

	for scanner.Scan() {
		names = append(names, readName(scanner.Text()))
	}

	return names
}

func readName(lineContent string) Person {
	var fname, lname string
	_, err := fmt.Sscanf(lineContent, "%s %s", &fname, &lname)
	if err != nil {
		panic(err)
	}
	return Person{fname, lname}
}

func printNames(names []Person) {
	for _, name := range names {
		fmt.Printf("first name: %20s   last name: %20s\n", name.fname, name.lname)
	}
}
