package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var path = "sample.txt" //path to problem input
var split = 0

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func extendBeam(beam []int, manifold []string, line int) []int {
	new_beam := make([]int, len(beam))

	for i, x := range beam {
		if x == 1 {
			if manifold[i] == "^" {
				new_beam[i+1] = 1
				new_beam[i-1] = 1
				split += 1
			} else {
				new_beam[i] = 1
			}
		}
	}
	return new_beam
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
	// var count int = 0
	matrix := [][]string{}

	for scanner.Scan() {
		// This is the "passing or at zero" variable
		matrix = append(matrix, strings.Split(scanner.Text(), ""))
	}
	// temp := len(matrix[0])
	beams := make([]int, len(matrix[0]))
	for i, x := range matrix[0] {
		if x == "S" {
			beams[i] = 1
		} else {
			beams[i] = 0
		}
	}
	for i, l := range matrix {
		if i == 0 {
			continue
		}
		beams = extendBeam(beams, l, i)
		// fmt.Println(beams)
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", split)

}
