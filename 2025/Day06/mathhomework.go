package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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
	problems := [][]int{}
	operand := []string{}
	const part1 = false

	if part1 == true {
		for scanner.Scan() {
			temp := strings.Split(scanner.Text(), " ")
			num := []string{}
			// Separate out the blanks, but otherwise separate into columns
			for _, n := range temp {
				if n != "" {
					num = append(num, n)
				}
			}

			for i, x := range num {
				value, err := strconv.Atoi(x)
				// convert strings to numbers... unless they're not numbers and then hope they're instructions
				if err != nil {
					operand = append(operand, x)
					continue
				} else {
					if len(problems) <= i {
						// If that slice hasn't been added, do so
						problems = append(problems, []int{})
					}
					problems[i] = append(problems[i], value)
				}
			}
		}

		// Solve all the problems
		for i, x := range problems {
			answer := 0
			for _, n := range x {
				if operand[i] == "*" {
					if answer != 0 { //ugh zeros, yo
						answer = answer * n
					} else {
						answer = n
					}
				} else {
					answer += n
				}
			}
			count += answer
		}
	} else { //PART II
		lines := []string{}
		for scanner.Scan() {
			temp := scanner.Text()
			// fmt.Println(temp)
			lines = append(lines, temp)
		}
		last := (lines[len(lines)-1])
		lines = slices.Delete(lines, len(lines)-1, len(lines)) //remove c
		c := []int{}
		for i, n := range last {
			if n != ' ' {
				c = append(c, i)
			}
		}

		// Ok... operands now contains a list of the starting position for all columns.
		// What next?
		// Columns can be defined by c[n] - c[n+1] until the last one
		max := len(c)
		for i := range c {
			problem := []string{}
			for _, l := range lines {
				num := ""
				if i+1 == max {
					num = (l[c[i]:])
				} else {
					num = (l[c[i] : c[i+1]-1]) //minus one needed to avoid the " "
				}
				if len(problem) != len(num) {
					problem = make([]string, len(num)) // Need to initiate the slice first pass only and make it the correct size
				}

				for i, n := range num {
					if n != ' ' {
						problem[i] = problem[i] + string(n)
					}
				}
			}

			answer := 0
			for _, n := range problem {
				number, _ := strconv.Atoi(n)
				if last[c[i]] == '*' {
					if answer != 0 {
						answer = answer * number
					} else {
						answer = number
					}

				} else {
					answer += number
				}
			}
			fmt.Println(problem, answer)
			count += answer
		}
	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", count)

}
