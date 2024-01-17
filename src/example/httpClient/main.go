package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	res, err := http.Get("https://hossein-moazami.ir")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	fmt.Println("Response status:", res.Status)
	fmt.Println("Response TLS:", res.TLS)

	scanner := bufio.NewScanner(res.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
