package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var path = "input.txt" //path to problem input
const bigly = 99999999999999999999.9
const part1 = false

// Function to handle errors
func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

type Smol struct {
	x   int
	y   int
	min float64
}

func findMin(line []float64, y int) Smol {
	var sm Smol
	sm.min = bigly
	for i, n := range line {
		if n == 0 {
			break
		} else if n > 0 && n < sm.min {
			sm.min = n
			// sm.coords = []int{i, y}
			sm.x = i
			sm.y = y
		}
	}
	return sm
}

func findSmolest(s []Smol) Smol {
	var sm Smol
	sm.min = 99999999999999999999.9
	for _, n := range s {
		if n.min < sm.min {
			sm.min = n.min
			sm.x = n.x
			sm.y = n.y
		}
	}
	return sm
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
	junction_boxes := [][]int{}
	pairs := [][]int{}

	for scanner.Scan() {
		// This is the "passing or at zero" variable
		temp := strings.Split(scanner.Text(), ",")
		coords := make([]int, 3)
		for i, x := range temp {
			coords[i], _ = strconv.Atoi(x)
		}
		junction_boxes = append(junction_boxes, coords)
		// fmt.Println(coords)

	}

	distances := make([][]float64, len(junction_boxes))
	distances_int := make([][]int, len(junction_boxes))
	smol := make([]Smol, len(junction_boxes))

	for y, j := range junction_boxes {
		if len(distances[y]) == 0 {
			distances[y] = make([]float64, len(junction_boxes))
			distances_int[y] = make([]int, len(junction_boxes))
		}
		for x, k := range junction_boxes {
			x_sq := (k[0] - j[0]) * (k[0] - j[0])
			y_sq := (k[1] - j[1]) * (k[1] - j[1])
			z_sq := (k[2] - j[2]) * (k[2] - j[2])
			sum := x_sq + y_sq + z_sq
			dist := math.Sqrt(float64(sum))
			distances[y][x] = dist
			distances_int[y][x] = int(dist)
		}
		smol[y] = findMin(distances[y], y)
		pairs = append(pairs, []int{y})
		// fmt.Println(distances_int[y], smol[y])
	}

	limit := 1000

	if !part1 {
		limit = limit * limit
	}

	var escape Smol

	// At this point we have all the junction boxes in their own "pairs" (solo, yo)
	// Now we need to iterate through them, shortest connection first, and connect them
	for i := 0; i < limit; i++ {

		next := findSmolest(smol)                         // Find the smallest of the small
		distances[next.y][next.x] = bigly                 // Delete it out of the list
		smol[next.y] = findMin(distances[next.y], next.y) // Reprocess that line
		// fmt.Printf("Shortest pair is x=%v y=%v, aka %v and %v ", next.x, next.y, junction_boxes[next.x], junction_boxes[next.y])

		// Check to see what pairs each is already in
		xy := []int{-1, -1}
		for n := range pairs {
			if slices.Contains(pairs[n], next.y) {
				xy[1] = n
			}
			if slices.Contains(pairs[n], next.x) {
				xy[0] = n
			}
		}

		// If they're already connected, do nothing, otherwise merge and delete
		if xy[0] == xy[1] {
			//i-- //already connected
			// fmt.Printf("which are already connected in group %v\n", xy[0])
			continue
		} else { //need to merge pairs?
			pairs[xy[0]] = append(pairs[xy[0]], pairs[xy[1]]...) //add one set to the other
			// fmt.Printf("which x is in pair %v and y in in pair %v, which will be merged and now size %v!\n", xy[0], xy[1], len(pairs[xy[0]]))
			pairs = slices.Delete(pairs, xy[1], xy[1]+1)

			if len(pairs) == 1 && i > 10 {
				escape.x = next.x
				escape.y = next.y
				break
			}
		}

	}

	if part1 {
		halp := []int{}
		for _, p := range pairs {
			halp = append(halp, len(p))
		}
		sort.Ints(halp)
		fmt.Println(halp)
		// And we're done!
	} else {
		fmt.Printf("The final two x coords are %v and %v, which multiply to %v, which is the answer to part 2.", junction_boxes[escape.x][0], junction_boxes[escape.y][0], junction_boxes[escape.x][0]*junction_boxes[escape.y][0])
	}

}
