//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Coordinate struct {
	x int
	y int
}
type Rectangle struct {
	a Coordinate //top-left
	b Coordinate //bottom-right
}

func width(rect Rectangle) int {
	return rect.b.x - rect.a.y
}

func length(rect Rectangle) int {
	return rect.a.y - rect.b.y
}

func calcArea(rect Rectangle) int {
	return width(rect) * length(rect)
}

func calcPerimeter(rect Rectangle) int {
	return (width(rect) * 2) + (length(rect) * 2)
}

func printTerminal(rectangle Rectangle) {
	fmt.Println("Area:", calcArea(rectangle))
	fmt.Println("Perimeter:", calcPerimeter(rectangle))

}
func main() {
	myRectangle := Rectangle{a: Coordinate{0, 7}, b: Coordinate{10, 0}}
	printTerminal(myRectangle)

	myRectangle.a.y *= 2
	myRectangle.b.x *= 2
	printTerminal(myRectangle)
}
