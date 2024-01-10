package main

import "fmt"

type Op func(int, int) int

func compute(lhs, rhs int, op Op) int {
	return op(lhs, rhs)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	mul := func(a, b int) int {
		return a * b
	}
	fmt.Println("7 * 8 =", compute(7, 8, mul))
	fmt.Println("---------------------")

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("---------------------")

	newInt := intSeq()
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println(newInt())
	fmt.Println("---------------------")

	fmt.Println(nextInt())

}
