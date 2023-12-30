//--Summary:
//  Create a program to read a list of numbers from multiple files,
//  sum the total of each file, then sum all the totals.
//
//--Requirements:
//* Sum the numbers in each file noted in the main() function
//* Add each sum together to get a grand total for all files
//  - Print the grand total to the terminal
//* Launch a goroutine for each file
//* Report any errors to the terminal
//
//--Notes:
//* This program will need to be ran from the `lectures/exercise/goroutines`
//  directory:
//    cd lectures/exercise/goroutines
//    go run goroutines
//* The grand total for the files is 4103109
//* The data files intentionally contain invalid entries
//* stdlib packages that will come in handy:
//  - strconv: parse the numbers into integers
//  - bufio: read each line in a file
//  - os: open files
//  - io: io.EOF will indicate the end of a file
//  - time: pause the program to wait for the goroutines to finish

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Sum func(digit int64)

func readFile(filesName []string, sum Sum) {
	for _, v := range filesName {
		file, err := os.Open(v)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			digit, err := strconv.ParseInt(scanner.Text(), 10, 32)
			if err != nil {
				log.Fatal(err)
			}
			// sum(digit)
			fmt.Println(digit)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func sumFile(rd bufio.Reader) int {
	sum := 0
	for {
		line, err := rd.ReadString('\n')
		if err == io.EOF {
			return sum
		}
		if err != nil {
			log.Fatal(err)
		}
		num, err := strconv.Atoi(line[:len(line)-1])
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
}
func main() {
	files := []string{"num1.txt", "num2.txt", "num3.txt", "num4.txt", "num5.txt"}

	totalSum := 0
	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			log.Fatal(err)
			return
		}
		defer file.Close()

		rd := bufio.NewReader(file)

		calculate := func() {
			fileSum := sumFile(*rd)
			totalSum += fileSum
		}
		go calculate()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(totalSum)

	// var total int64
	// sumDigit := func(digit int64) {
	// 	total += digit
	// }

	// readFile(files, sumDigit)
	// fmt.Println("Sum is ", total)
}
