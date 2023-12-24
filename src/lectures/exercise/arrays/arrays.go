//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	name  string
	count int
	price int
}

func add(name string, count int, price int) Product {
	return Product{name, count, price}
}

func printTerminal(items [4]Product) {
	fmt.Println("# R", "\tName\tCount\tPrice")
	fmt.Println("---------------------------------")
	var totalCount, totalPrice int
	for i := 0; i < len(items); i++ {
		element := items[i]

		if element.name != "" {
			totalCount += element.count
			totalPrice += element.price
			fmt.Println("#", i+1, element.name, "\t", element.count, "\t", element.price)
		}
	}
	fmt.Println("---------------------------------")
	fmt.Println("#", "\t\t", totalCount, "\t", totalPrice)
}

func main() {
	shoppingList := [4]Product{}

	shoppingList[0] = add("Banana", 1, 7)
	shoppingList[1] = add("Orange", 5, 5)
	shoppingList[2] = add("Tomato", 10, 3)
	shoppingList[3] = add("Peach", 7, 15)
	printTerminal(shoppingList)
}
