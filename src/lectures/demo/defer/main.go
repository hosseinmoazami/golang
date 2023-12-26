package main

import (
	"fmt"
	"os"
)

func main() {
	readNameFile("lectures/demo/defer/names")
}

func readNameFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	buffer := make([]byte, 0, 30)
	bytes, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%c\n", bytes)
}
