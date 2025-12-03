package main

import (
	"bufio"
	"fmt"
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
		biggest := 0

		// This is stupid and I hate it
		var battery []int
		for _, j := range temp {
			temp2, _ := strconv.Atoi(j)
			battery = append(battery, temp2)
		}
		for x, jolt := range battery {
			for p := x + 1; p < bank_size; p++ {
				bat := (jolt * 10) + battery[p]
				if bat > biggest {
					biggest = bat
				}
			}
		}
		// fmt.Println(biggest)
		count += biggest
	}
	fmt.Println(count)
}
