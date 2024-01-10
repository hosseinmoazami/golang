package main

import "fmt"

func main() {
	// pizza := []int{3, 2, 5, 23, 342, 12, 665, 45}
	pizza := make([]int, 6, 7)
	fmt.Printf("Cap: %v\n Len: %v\n", cap(pizza), len(pizza))
	for i := 0; i < len(pizza); i++ {
		element := pizza[i]
		fmt.Println(element)
	}
}
