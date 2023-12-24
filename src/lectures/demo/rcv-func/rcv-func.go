package main

import "fmt"

type Space struct {
	occupied bool
}
type ParkingSpace struct {
	spaces []Space
}

func occupySpace(lot *ParkingSpace, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingSpace) occupySpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingSpace) vacateSpace(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {
	lot := ParkingSpace{spaces: make([]Space, 4)}
	fmt.Println("Initial", lot)

	lot.occupySpace(1)
	occupySpace(&lot, 2)
	fmt.Println("After Occupy", lot)

	lot.vacateSpace(2)
	fmt.Println("After Vacate", lot)
}
