package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

const TOTALROWS = 128
const TOTALCOLUMNS = 8

func main() {
	f, _ := os.Open("day5_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	maxID := 0
	for _, line := range lines {
		row := binarySearch(line[:7], TOTALROWS-1)       // subtract 1 because indexed by 0
		column := binarySearch(line[7:], TOTALCOLUMNS-1) // subtract 1 because indexed by 0

		id := row*8 + column
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println(maxID)

}

func binarySearch(in string, upperBound int) int {
	lower := 0
	upper := upperBound
	for _, c := range in {
		diff := upper - lower // difference between high and low
		switch string(c) {
		case "F", "L":
			upper = upper - int(math.Ceil(float64(diff)/2.0))
		case "B", "R":
			lower = lower + int(math.Ceil(float64(diff)/2.0))
		}
	}

	if len(in) > 3 {
		return lower
	}
	return upper
}
