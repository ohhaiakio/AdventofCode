package main

import (
	"bufio"
	"fmt"
	"os"
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

func sumslice(slice [][]int, x int, y int, s int) int {
	sum := 0
	for yy := y - s; yy <= y+s; yy++ {
		for xx := x - s; xx <= x+s; xx++ {
			sum += slice[yy][xx]
		}
	}
	// test := slice[y-1 : y+2][x-1 : x+2]
	// fmt.Println(test)
	return (sum - 1)
}

func scan_Grid(grid [][]int) [][]int {
	moveable := [][]int{}

	xlen := len(grid[0])
	ylen := len(grid)
	for y := 1; y < ylen-1; y++ {
		for x := 1; x < xlen-1; x++ {
			//now we have the x,y of the paper we want to look at, we need to look at all adjacent spaces
			if grid[y][x] == 1 {
				adjacent := sumslice(grid, x, y, 1)
				if adjacent < 4 {
					moveable = append(moveable, []int{x, y})
				}
			}
		}
	}
	return moveable
}

func move_Paper(coords [][]int, grid [][]int) [][]int {
	for _, c := range coords {
		grid[c[1]][c[0]] = 0
	}
	return grid
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
	var total int = 0
	grid := [][]int{}

	for scanner.Scan() {
		txt := scanner.Text()
		// size := len(txt)
		temp := strings.Split(txt, "")
		const paper = "@"
		row := []int{0}
		for _, x := range temp {
			if x == paper {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		row = append(row, 0)
		grid = append(grid, row)
		// fmt.Println(row)
	}
	//prepend zeros to the first and last positions to make life easier
	grid = append(grid, make([]int, len(grid[0])))
	grid = append([][]int{make([]int, len(grid[0]))}, grid...)

	count := 99
	for count > 0 {
		moveable := scan_Grid(grid)
		grid = move_Paper(moveable, grid)
		count = len(moveable)
		total += count
		// fmt.Println("Moved", count, "rolls of paper!")
	}
	fmt.Println("All rolls have been moved! Moved", total, "rolls of paper in total!")

	// moveable := scan_Grid(grid)
	// grid = move_Paper(moveable, grid)
	// moveable = scan_Grid(grid)
	// fmt.Println(len(moveable))
}
