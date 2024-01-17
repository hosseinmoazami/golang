package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("path", "/etc/resolve.conf")
	fmt.Println("Path", os.Getenv("path"))
	fmt.Println("Name", os.Getenv("Name"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
