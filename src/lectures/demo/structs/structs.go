package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {
	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	fatima := Passenger{"Fatima", 2, true}
	fmt.Println(fatima)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 3}
		ella = Passenger{Name: "Ella", TicketNumber: 4}
	)
	fmt.Println(bill, ella)

	var hossein Passenger
	hossein.Name = "Hossein"
	hossein.TicketNumber = 5
	fmt.Println(hossein)

	casey.Boarded = true
	fatima.Boarded = true
	if casey.Boarded {
		fmt.Println("Casey has boarded on the Bus")
	}
	if fatima.Boarded {
		fmt.Println("Fatima has boarded on the Bus")
	}

	hossein.Boarded = true
	bus := Bus{hossein}
	fmt.Println(bus)
	fmt.Println(bus.FrontSeat.Name, "is in the front seat")
}
