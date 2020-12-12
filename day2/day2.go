package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day2_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	validCount := 0
	for _, line := range lines {
		if processLine(line) {
			validCount++
		}
	}
	fmt.Println(validCount)

}

func processLine(in string) bool {
	tup := strings.Split(in, " ")

	countRange := tup[0]
	rangeSplit := strings.Split(countRange, "-")
	lower, _ := strconv.Atoi(rangeSplit[0])
	upper, _ := strconv.Atoi(rangeSplit[1])

	letter := strings.ReplaceAll(tup[1], ":", "")

	password := tup[2]

	return isValid(lower, upper, letter, password)

}

func isValid(lower, upper int, letter, password string) bool {
	c := strings.Count(password, letter)
	if c >= lower && c <= upper {
		return true
	}
	return false
}
