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

	var seats [][]int = make([][]int, TOTALROWS)
	for s := range seats {
		column := make([]int, TOTALCOLUMNS)
		seats[s] = column
	}

	for _, line := range lines {
		row := binarySearch(line[:7], TOTALROWS-1)       // subtract 1 because indexed by 0
		column := binarySearch(line[7:], TOTALCOLUMNS-1) // subtract 1 because indexed by 0
		seats[row][column] = -1
	}

	for i, columns := range seats {
		for j := range columns {
			if seats[i][j] == 0 && i > 5 && i < 124 { // make sure not at the "very front or back" of plane
				fmt.Println(i*8 + j)
			}
		}
	}

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

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
