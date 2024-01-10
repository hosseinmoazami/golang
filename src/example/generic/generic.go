package main

import "fmt"

// non-generic
func sumOfInt(input []int) int {
	inc := 0
	for _, val := range input {
		inc += val
	}
	return inc
}

// generic
func sumOfFloat[T float32](input []T) T {
	var inc T
	for _, val := range input {
		inc += val
	}
	return inc
}

type Number interface {
	int | float32
}

func sumOfAny[T Number](input []T) T {
	var inc T
	for _, val := range input {
		inc += val
	}
	return inc
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []float32{1.0, 2.0, 3.0, 4.0}

	fmt.Printf("sum of int: %v\n", sumOfInt(a))
	fmt.Printf("sum of float32: %0.2f\n", sumOfFloat(b))
	fmt.Println("_____________________________")
	fmt.Printf("sum of Any(int): %v\n", sumOfAny(a))
	fmt.Printf("sum of Any(float): %0.2f\n", sumOfAny(b))
}
