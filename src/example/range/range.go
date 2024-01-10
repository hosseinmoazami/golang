package main

import "fmt"

func main() {
	myMap := make(map[string]string)

	myMap["hair_color"] = "black"
	myMap["eye_color"] = "green"
	myMap["height"] = "tall"
	myMap["skin_color"] = "white"
	myMap["adult"] = "middle"

	for key, value := range myMap {
		fmt.Printf("%v=>%v\n", key, value)
	}
}
