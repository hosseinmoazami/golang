package main

import (
	"fmt"
)

type Person struct {
	ID        uint32
	firstName string
	lastName  string
	age       uint
	phone     string
}

func add(ID uint32, firstName string, lastName string, age uint, phone string) Person {
	return Person{
		ID:        ID,
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		phone:     phone,
	}
}

func main() {
	persons := make([]Person, 0)

	persons = append(persons, add(1, "Hossein", "Loe", 27, "985521"))
	persons = append(persons, add(2, "Lina", "Kube", 54, "982214"))

	for _, person := range persons {
		fmt.Println(person)

	}
}
