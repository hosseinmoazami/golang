package main

import (
	"fmt"
	"time"
	"unicode"
)

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitalized []rune

	capIt := func(r rune) {
		capitalized = append(capitalized, unicode.ToUpper(r))
		fmt.Println()
		fmt.Printf("%c Done!\n", r)
		fmt.Println()
	}

	for i := 0; i < len(data); i++ {
		go capIt(data[i])
		go hugeCalculation("Hello mamesita mamaesima caltuera")
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("Capitalized: %c\n", capitalized)
}

func hugeCalculation(text string) {
	for k, v := range text {
		fmt.Printf("[%c] is the char [%v] in text.\n", v, k)
	}
}
