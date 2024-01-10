package main

import "fmt"

func Sum(a, b int) int {
	sum := a + b
	return sum
}

func comet(sign string, list ...int) {
	for i := 0; i < len(list); i++ {
		element := list[i]
		for j := 0; j < element; j++ {
			fmt.Printf("%v", sign)
		}
		fmt.Println()
	}
}

func main() {
	// fmt.Printf("1 + 2 = %v\n", Sum(1, 2))
	// fmt.Printf("8 + 6 = %v\n", Sum(8, 6))
	// fmt.Printf("4 + 7 = %v\n", Sum(4, 7))
	// fmt.Printf("17 + 42 = %v\n", Sum(17, 42))

	oranos := []int{1, 2, 3}
	petos := []int{5, 12, 32}
	xienos := []int{7, 7, 7, 7}
	famous := []int{1, 8, 14}

	comet("*", oranos...)
	comet("*", petos...)
	comet("*", xienos...)
	comet("*", famous...)
}
