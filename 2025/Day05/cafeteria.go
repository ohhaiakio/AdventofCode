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

func inRange(id int, fresh [][]int) bool {
	for _, r := range fresh {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

func inRange2(start int, finish int, fresh [][]int) []int {
	for _, r := range fresh {
		if start >= r[0] && start <= r[1] {
			start = r[1] + 1
		}
		if finish >= r[0] && finish <= r[1] {
			finish = r[0] - 1
		}
	}
	if start > finish {
		start = -99
		finish = -99
	}
	return []int{start, finish}
}

func inRange3(id int, fresh [][]int) int {
	c := 0
	hits := []int{}
	if id == -99 {
		return 1
	}
	for i, r := range fresh {
		if id >= r[0] && id <= r[1] {
			hits = append(hits, i)
			c++
		}
	}
	// if c != 1 {
	// 	fmt.Println(id, c, hits)
	// }
	return c
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
	part2 := 0

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
			new_range := inRange2(start, stop, fresh)
			new := new_range[1] - new_range[0] + 1
			if new < 0 {
				new = 0
			} else {
				fresh = append(fresh, new_range)
			}

		} else {
			id, _ := strconv.Atoi(txt)
			if inRange(id, fresh) {
				count += 1
			}
		}
	}
	// And we're done!
	fmt.Printf("\nThere were %v fresh ingredients in the first part.\n", count)
	part2 = 0
	for x := range fresh {
		c := 0
		c += inRange3(fresh[x][0], fresh)
		c += inRange3(fresh[x][1], fresh)
		new := 0
		if c != 2 {
			// fmt.Println("Found a problem!", x, c)
			fresh[x] = inRange2(fresh[x][0], fresh[x][1], fresh)
		}
		new = fresh[x][1] - fresh[x][0] + 1
		if new > 0 {
			part2 += new
		}
	}
	for x := range fresh {
		c := 0
		c += inRange3(fresh[x][0], fresh)
		c += inRange3(fresh[x][1], fresh)
		if c != 2 {
			fmt.Println("Found a problem!", x, c)
			// fresh[x] = inRange2(fresh[x][0], fresh[x][1], fresh)
		}
	}

	fmt.Printf("\nThere are %v total fresh ingredients available.\n", part2)
}
