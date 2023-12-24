package main

import "fmt"

func main() {

	shoppingList := make(map[string]int)

	shoppingList["eggs"] = 10
	shoppingList["milk"] = 1
	shoppingList["oil"] = 3
	shoppingList["bread"] = 5
	shoppingList["wine"] = 6

	shoppingList["eggs"] += 2
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted, new list:", shoppingList)

	fmt.Println("Need Cereal?")

	cereal, found := shoppingList["cereal"]
	if !found {
		fmt.Println("Nope")
	} else {
		fmt.Println("Yup,", cereal, "boxes")
	}

	totalItems := 0
	for _, amount := range shoppingList {
		totalItems += amount
	}

	fmt.Println("Total Items is:", totalItems, "in the shopping list")
}
