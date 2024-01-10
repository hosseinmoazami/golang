package main

import "fmt"

func main() {
	beuget := make(map[string]int)

	beuget["item"] = 3
	beuget["price"] = 2500
	beuget["log"] = 34
	beuget["poto"] = 0

	fmt.Println(beuget)

}
