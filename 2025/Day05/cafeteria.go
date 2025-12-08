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
	old_start := start
	old_finish := finish

	for _, r := range fresh {
		if old_start == r[0] || old_finish == r[1] {
			// fmt.Printf("Something weird happened. %v %v %v\n", r[0], r[1])
			return []int{-999, -99}
		}
		if start >= r[0] && start <= r[1] {
			start = r[1] + 1
		}
		if finish >= r[0] && finish <= r[1] {
			finish = r[0] - 1
		}
	}
	if start > finish {
		start = -999
		finish = -99
	}
	return []int{start, finish}
}

func inRange3(id int, fresh [][]int) int {
	c := 0
	// hits := []int{}
	if id == -99 || id == -999 {
		return 1
	}
	for _, r := range fresh {
		if id >= r[0] && id <= r[1] {
			// hits = append(hits, i)
			c++
		}
	}
	// if c != 1 {
	// 	fmt.Println(id, c, hits)
	// }
	return c
}

func detectRange(start int, stop int, fresh [][]int) int {
	c := 0
	c += inRange3(start, fresh)
	c += inRange3(stop, fresh)
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
	hello_fresh := [][]int{}
	for x := range fresh {

		hello_fresh = append(hello_fresh, inRange2(fresh[x][0], fresh[x][1], hello_fresh))

		for detectRange(fresh[x][0], fresh[x][1], fresh) > 2 {
			old := fresh[x]
			fresh[x] = []int{-777, -777}
			// fresh[x] = inRange2(fresh[x][0], fresh[x][1], fresh)
			fresh[x] = inRange2(old[0], old[1], fresh)
		}

		if fresh[x][1] > 0 {
			calc := fresh[x][1] - fresh[x][0] + 1
			if calc > 0 {
				part2 += calc
			}
		}

		// fmt.Printf("\nStart: %v\tFinish: %v\tCalc: %v\n", fresh[x][0], fresh[x][1], calc)

	}

	for x := range fresh {
		if fresh[x][1] > 0 {
			calc := fresh[x][1] - fresh[x][0] + 1
			if calc > 0 {
				part2 += calc
			}
		}
	}
	fmt.Printf("\nThere are %v total fresh ingredients available.\n", part2)
	part2 = 0
	for x := range hello_fresh {
		if hello_fresh[x][1] > 0 {
			calc := hello_fresh[x][1] - hello_fresh[x][0] + 1
			if calc > 0 {
				part2 += calc
			}
		}
	}

	fmt.Printf("\nThere are %v total fresh ingredients available.\n", part2)
}
