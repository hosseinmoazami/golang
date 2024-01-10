package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (rect *Rectangle) area() int {
	return rect.height * rect.width
}

func (rect Rectangle) perim() int {
	return rect.height*2 + rect.width*2
}

func main() {
	rectangle := Rectangle{width: 10, height: 5}

	r := &rectangle

	fmt.Println("Area:", rectangle.area())
	fmt.Println("perim:", r.perim())
}
