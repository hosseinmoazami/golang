package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	all := append(a, b...)
	answer := sum(all...)
	fmt.Println(answer)

	d := sum(1, 2, 3, 4, 5, 6)
	fmt.Println(d)
}
