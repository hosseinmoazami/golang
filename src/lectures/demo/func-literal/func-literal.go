package main

import (
	"fmt"
	"strings"
)

type Op func(lhs, rhs int) int

func main() {
	customMsg(surround(), "hossein moazami")
	fmt.Println("2 + 2 =", compute(2, 2, add))
	fmt.Println("20 + 25 =", compute(2, 2, func(lhs, rhs int) int {
		return lhs - rhs
	}))

	mul := func(lhs, rhs int) int {
		fmt.Printf("Multiplying %v * %v = ", lhs, rhs)
		return lhs * rhs
	}
	fmt.Println(compute(5, 4, mul))

}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func compute(lhs, rhs int, op Op) int {
	fmt.Printf("Running a computation with %v & %v\n", lhs, rhs)
	return op(lhs, rhs)
}

func customMsg(fn func(m string), msg string) {
	msg = strings.ToUpper(msg)
	fn(msg)
}

func surround() func(msg string) {
	return func(msg string) {
		fmt.Printf("%.*s\n", len(msg), "--------------------------------")
		fmt.Println(msg)
		fmt.Printf("%.*s\n", len(msg), "--------------------------------")
	}
}
