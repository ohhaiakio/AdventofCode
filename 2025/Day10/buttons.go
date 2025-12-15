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

func pressButtons(buttons []int64, possibilites []int64, lights int64, presses int) int {
	// for buttons
	new_possibilities := []int64{}
	for _, p := range possibilites {
		for _, b := range buttons {
			press := p ^ b
			if press == lights {
				return presses + 1
			} else if !slices.Contains(new_possibilities, press) {
				new_possibilities = append(new_possibilities, press)
			}
		}
	}
	return pressButtons(buttons, new_possibilities, lights, presses+1)

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

	for scanner.Scan() {
		line := scanner.Text()
		butt_string := ""
		lit_string := ""
		jolt_string := ""

		// butt := 0
		// var lit_string string
		// lit_string := ""
		ref := 0

		for i, c := range line {
			if c == ']' {
				lit_string = line[1:i]
				ref = i
			} else if c == '{' {
				butt_string = line[ref+2 : i-1]
				jolt_string = line[i:]
				break
			}
		}
		fmt.Println(butt_string, lit_string, jolt_string)

		// Convert the lights to binary and then to ints
		bin_string := ""
		blank := []rune{}
		for _, x := range lit_string {
			if x == '#' {
				bin_string += "1"
			} else {
				bin_string += "0"
			}
			blank = append(blank, '0')
		}
		l, _ := strconv.ParseInt(bin_string, 2, 64)

		// fmt.Println(butt, l)

		// Convert buttons
		butt_string = strings.ReplaceAll(butt_string, "(", "")
		butt_string = strings.ReplaceAll(butt_string, ")", "")
		temp := strings.Split(butt_string, " ")
		temp2 := []int64{}
		for _, x := range temp {
			omgwhy := make([]rune, len(blank))
			copy(omgwhy, blank)
			t := strings.Split(x, ",")
			// omgwhy := []int{}
			// omgwhy := "00000000000"
			for _, y := range t {
				i, _ := strconv.Atoi(y)
				// wtf := len(omgwhy) - 1 - i
				wtf := i
				omgwhy[wtf] = '1'
			}
			// fmt.Println(string(omgwhy))
			buttons_for_real, _ := strconv.ParseInt(string(omgwhy), 2, 64)
			temp2 = append(temp2, buttons_for_real)
		}
		fmt.Println(temp2)

		presses := pressButtons(temp2, []int64{0}, l, 0)
		count += presses
		// fmt.Println("Oh hey we found a thing and it was ", presses)

	}

	// And we're done!
	fmt.Printf("\nI think we're finished and the count was %v.\n", count)

}
