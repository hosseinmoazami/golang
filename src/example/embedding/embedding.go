package main

import "fmt"

type Base struct {
	num int
}

func (base Base) describe() string {
	return fmt.Sprintf("base with num=%v", base.num)
}

type Container struct {
	Base
	str string
}

func main() {
	c := Container{
		Base: Base{
			num: 43,
		},
		str: "Hello",
	}

	fmt.Println(c.num)
	fmt.Println(c.Base.num)
	fmt.Println(c.str)
	fmt.Println(c.describe())

	type describer interface {
		describe() string
	}

	var d describer = c
	fmt.Println("Describer:", d.describe())
}
