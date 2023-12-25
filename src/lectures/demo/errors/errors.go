package main

import (
	"errors"
	"fmt"
)

type Stuff struct {
	values []int
	names  []string
	code   map[int]string
}

func (s *Stuff) Get(index int) (int, error) {
	if index > len(s.values) {
		return 0, errors.New(fmt.Sprintf("no element at index %v", index))
	} else {
		return s.values[index], nil
	}
}

func main() {
	stuff := Stuff{values: []int{0, 1}, names: []string{"hossein", "ali"}, code: map[int]string{1: "A", 2: "B"}}
	fmt.Println(stuff)

	value, err := stuff.Get(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("value is", value)
	}
}
