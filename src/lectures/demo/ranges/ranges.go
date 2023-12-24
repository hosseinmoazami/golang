package main

import "fmt"

func main() {
	slice := []string{"Hello", "World", "!"}

	for index, value := range slice {
		fmt.Println(index, value, ":")

		for _, ch := range value {
			fmt.Printf("    %q\n", ch)
		}
	}
}
