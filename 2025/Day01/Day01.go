package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var path = "input.txt" //path to problem input

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func Abs(i int) int {
	if i < 0 {
		return (100 + i)
	}
	return i
}

func main() {
	//Open input file
	var file, err = os.Open(path)
	if isError(err) {
		return
	}

	//Load all lines in file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// //initialize output variable
	var count int = 0

	//Initialize variables (might not be needed?)
	var dial int = 50

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		dir := scanner.Text()[:1]
		num, _ := strconv.Atoi(scanner.Text()[1:])
		// fmt.Println(dir, num)
		if dir == "L" {
			dial = Abs(dial-num) % 100

		} else {
			dial = (dial + num) % 100
		}
		if dial == 0 {
			count += 1
		}
		// fmt.Println(dial)
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", count)

}
