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
		// This is the "passing or at zero" variable
		wrap := 0

		// Grab instructions and save old data
		previous := dial
		dir := scanner.Text()[:1]
		num, err := strconv.Atoi(scanner.Text()[1:])
		if err != nil {
			fmt.Println(err)
			break
		}

		// Left or right?
		if dir == "L" {
			dial = dial - num
			if dial < 0 {
				wrap = (-dial / 100) + 1     //determine how many times we pass zero
				dial = ((wrap) * 100) + dial //spicy absolute value
				if previous == 0 {           //needed to not double count passing zero
					wrap += -1
				}
				dial = dial % 100 //edge case where we accidentally make a 0 a 100 but it's fine i'm sure
			} else if dial == 0 { // another edge where this zero doesn't get counted
				wrap += 1
			}
		} else { //R
			dial = (dial + num)
			wrap += dial / 100
			dial = dial % 100
		}

		count += wrap
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", count)

}
