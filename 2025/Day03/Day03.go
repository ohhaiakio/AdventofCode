package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

// Function recursively tries to find the position of the highest digit value left
// Function doesn't consider values too far into the string to be useful (not enough digits left)
func find_nine(b []int, start int, stop int, size int, nine int) int {
	for x := start; x < (size - stop); x++ {
		if b[x] == nine {
			return x
		}
	}
	return find_nine(b, start, stop, size, nine-1)
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
	// var dial int = 50

	for scanner.Scan() {
		line := scanner.Text()
		bank_size := len(line)
		temp := strings.Split(line, "")

		// This is stupid and I hate it
		var battery []int
		for _, j := range temp {
			temp2, _ := strconv.Atoi(j)
			battery = append(battery, temp2)
		}

		// PART 1
		// biggest := 0
		// for x, jolt := range battery {
		// 	for p := x + 1; p < bank_size; p++ {
		// 		bat := (jolt * 10) + battery[p]
		// 		if bat > biggest {
		// 			biggest = bat
		// 		}
		// 	}
		// }

		//Part 2
		baby_batt := 0 //actually just the storage for the new battery
		start := 0     //starting point for digit scan - stores last found digit value
		stop := 11     //stopping point for digit scan - reduces as baby_batt grows
		for stop >= 0 {
			p := find_nine(battery, start, stop, bank_size, 9)
			baby_batt = baby_batt + int((math.Pow10(stop)))*battery[p]
			// fmt.Println(baby_batt, start, stop, p)
			stop--
			start = p + 1
		}

		count += baby_batt
	}
	fmt.Println(count)
}
