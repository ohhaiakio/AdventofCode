package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var path = "input.txt" //path to problem input
var count = 0
var part2 = false

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func navigate(m map[string][]string, start string, path []string, c int) int {
	for _, next := range m[start] {
		if next == "out" {
			if part2 {
				fmt.Println(path, count+1)
				if slices.Contains(path, "dac") && slices.Contains(path, "fft") {
					// if slices.Index(path, "fft") > slices.Index(path, "dac") {
					fmt.Println(path, count+1)
					count++
					// }
				}
			} else {
				// fmt.Println(path, count+1)
				count++
			}

		} else if !slices.Contains(path, next) {
			navigate(m, next, append(path, next), count)
		}
	}
	return count
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
	m := make(map[string][]string)

	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		m[temp[0][:3]] = temp[1:]
	}
	navigate(m, "you", []string{"you"}, 0)

	// And we're done!
	fmt.Printf("\nI think we're finished Part 1 and the count was %v.\n", count)

	count = 0
	part2 = true
	navigate(m, "svr", []string{"svr"}, 0)
	fmt.Printf("\nI think we're finished Part 2 and the count was %v.\n", count)

}
