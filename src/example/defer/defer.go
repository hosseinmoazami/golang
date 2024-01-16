package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f, "Hello This is a test message")

}

func createFile(name string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, data string) {
	fmt.Println("writing")
	fmt.Fprintln(f, data)
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
