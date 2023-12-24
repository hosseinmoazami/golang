//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

func printInfo(items []Product) {
	for i := 0; i < len(items); i++ {
		item := items[i]
		fmt.Println(item)
	}
}

func initialItem() []Product {
	shirt := Product{"Shirt", Active}
	pants := Product{"Pants", Active}
	purse := Product{"Purse", Active}
	watch := Product{"Watch", Active}

	items := []Product{shirt, pants, purse, watch}
	return items
}

type SecurityTag bool

const (
	Active   = true
	Inactive = false
)

type Product struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func deactivate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Product) {
	fmt.Println("Checking out...")
	for i := 0; i < len(items); i++ {
		deactivate(&items[i].tag)
	}
}

func checkIn(item *Product) {
	fmt.Println("Checking In...")
	activate(&item.tag)
}

func main() {
	items := initialItem()
	printInfo(items)
	fmt.Println("=================")
	deactivate(&items[0].tag)
	printInfo(items)
	fmt.Println("=================")
	checkout(items)
	checkIn(&items[2])
	printInfo(items)
	fmt.Println("=================")

}
