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

type Coords struct {
	x int
	y int
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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

	//Initialize variables (might not be needed?)
	red_tiles := []Coords{}

	for scanner.Scan() {
		// This is the "passing or at zero" variable
		temp := strings.Split(scanner.Text(), ",")
		var coords Coords
		coords.x, _ = strconv.Atoi(temp[0])
		coords.y, _ = strconv.Atoi(temp[1])
		red_tiles = append(red_tiles, coords)
		// fmt.Println(coords)
	}

	// squares := make([][]int, len(red_tiles))
	biggest := 0
	for _, a := range red_tiles {
		for _, b := range red_tiles {
			size := (Abs(a.x-b.x) + 1) * (Abs(a.y-b.y) + 1)
			if size > biggest {
				biggest = size
			}
		}
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", biggest)

}
