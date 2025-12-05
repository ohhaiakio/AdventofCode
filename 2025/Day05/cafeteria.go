package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var path = "sample.txt" //path to problem input

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func inRange(id int, fresh [][]int) bool {
	for _, r := range fresh {
		if id >= r[0] && id <= r[1] {
			fmt.Println(id)
			return true
		}
	}
	return false
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

	fresh := [][]int{}
	preamble := true

	for scanner.Scan() {
		// Grab instructions and save old data
		txt := scanner.Text()

		if txt == "" {
			preamble = false
			continue
		} else if preamble {
			temp := strings.Split(txt, "-")
			start, _ := strconv.Atoi(temp[0])
			stop, _ := strconv.Atoi(temp[1])
			fresh = append(fresh, []int{start, stop})
		} else {
			id, _ := strconv.Atoi(txt)
			if inRange(id, fresh) {
				count += 1
			}
		}
	}
	// And we're done!
	fmt.Printf("\nThere were %v fresh ingredients in the first part.\n", count)

	// PART 2

}
