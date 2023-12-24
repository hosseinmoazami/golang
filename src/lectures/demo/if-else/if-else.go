package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {
	quiz1, quiz2, quiz3 := 4, 7, 10
	var average float32 = average(quiz1, quiz2, quiz3)

	if quiz1 > quiz2 {
		fmt.Println("Q1 is grater than Q2")
	} else if quiz1 < quiz2 {
		fmt.Println("Q1 is less than Q2")
	} else {
		fmt.Println("Q1 and Q2 are equal")
	}

	if average > 7 {
		fmt.Println("Result is acceptable")
	} else {
		fmt.Println("You failed, your average is:", average)
	}
}
