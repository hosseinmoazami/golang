//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var favoriteColor string = "green"
	brithYear, age := 1996, 27
	// var (
	// 	first = 1
	// 	last  = 20
	// )
	var ageInDays int
	ageInDays = age * 365 * 24 * 60 * 60

	fmt.Println("My favorite Color is", favoriteColor)
	fmt.Println("I was born in", brithYear, "And I am", age, "Years old")
	fmt.Println("My age in days is", ageInDays)

}
