package main

import "fmt"

func average(nums [5]float32) float32 {
	var sum float32
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return float32(sum / 5)
}

func main() {
	var buckets [3]int
	scores := [5]float32{17.25, 16.5, 20, 10.25, 9.25}
	matrix := [2][2]int{{4, 4}, {5, 4}}
	fmt.Println(matrix)

	avg := average(scores)
	fmt.Printf("Average is: %v\n", avg)

	buckets[0] = 5
	buckets[1] = 7
	buckets[2] = 12

	fmt.Println(buckets)
	for i := 0; i < len(buckets); i++ {
		element := buckets[i]
		fmt.Print(element, ">")

		for j := 0; j < element; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
