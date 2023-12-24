package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func visitingLocations(indexs []int, slice []string) {
	fmt.Println()
	fmt.Println("--Visiting--")
	for i := 0; i < len(indexs); i++ {
		element := slice[i]
		fmt.Println("\t\t", element)
	}
}

func remainingLocations(index int, slice []string) {
	fmt.Println()
	fmt.Println("--Remaining--")
	remain := slice[index:]
	for i := 0; i < len(remain); i++ {
		element := remain[i]
		fmt.Println("\t\t", element)
	}
}

func main() {
	route := []string{"Grocery", "Deportment Store", "Salon"}
	printSlice("Route #1", route)

	route = append(route, "Home")
	printSlice("Route #2", route)

	indexs := []int{0, 1}
	visitingLocations(indexs, route)

	remainingLocations(len(indexs), route)
}
