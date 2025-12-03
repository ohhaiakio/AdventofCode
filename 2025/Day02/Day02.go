package main

import (
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

func part1(input []string) {
	count := 0

	for _, id := range input {
		r := strings.Split(id, "-")

		fin, _ := strconv.Atoi(r[1])
		// Iterate over all of the possible IDs and see if they're symmetrical
		for num, _ := strconv.Atoi(r[0]); num <= fin; num++ {
			num_st := strconv.Itoa(num)
			num_len := len(num_st)

			// If the number has an odd number of digits, we can ignore it
			if (num_len % 2) != 0 {
				continue //odd
			}

			// Otherwise, split and compare
			split := num_len / 2
			if num_st[:split] == num_st[split:] {
				count += num
			}
		}
	}
	fmt.Println("I think we've finished and the sum of all the invalid codes is:", count)
}

func part2(input []string) {
	count := 0

	for _, id := range input {
		r := strings.Split(id, "-")

		fin, _ := strconv.Atoi(r[1])
		// Iterate over all of the possible IDs and see if they're symmetrical
		for num, _ := strconv.Atoi(r[0]); num <= fin; num++ {
			num_st := strconv.Itoa(num)

		NumCheck:
			for i := 1; i < len(num_st); i++ {
				check := num_st[:i]
				mult := len(num_st[i:]) % (i)
				if mult != 0 {
					continue
				}
				div := len(num_st[i:]) / i

				for n := 0; n < div; n++ {
					// monster := num_st[i+(n*i) : i+i+(n*i)]
					if check == num_st[i+(n*i):i+i+(n*i)] {
						if n+1 == div {
							count += num
							fmt.Println(num)
							break NumCheck
						}
					} else {
						break
					}
				}
			}
		}
	}
	fmt.Println("I think we've finished and the sum of all the invalid codes is:", count)
}

func main() {
	//Open input file
	var file, err = os.ReadFile(path)
	if isError(err) {
		return
	}
	input := strings.Split(string(file), ",")

	part2(input)
}
