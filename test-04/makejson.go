package main

import "fmt"
import "encoding/json"

func main() {
	var name string
	var address string
	preJson := make(map[string]string)

	fmt.Println("Enter your name:")
	fmt.Scan(&name)

	fmt.Println("Enter your address:")
	fmt.Scan(&address)

	preJson["name"] = name
	preJson["address"] = address

	newJson, err := json.Marshal(preJson)
	if err != nil {
		fmt.Println("Could not convert to Json")
	} else {
		fmt.Println(string(newJson))
		fmt.Println(newJson)
	}
}
