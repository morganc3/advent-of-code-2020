package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day9_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	lines_ints := make([]int, len(lines))
	for i, e := range lines {
		int_val, _ := strconv.Atoi(e)
		lines_ints[i] = int_val
	}

	previousCheckCount := 25
	preamble := 25

	for i := range lines_ints {
		if i < preamble {
			continue
		}
		sumFound := findSumInPrevious(lines_ints[i-previousCheckCount:i], lines_ints[i])
		if !sumFound {
			fmt.Println(lines_ints[i])
			return
		}
	}

}

func findSumInPrevious(prev []int, sum int) bool {
	copyToSort := make([]int, len(prev))
	copy(copyToSort, prev)
	sort.Ints(copyToSort)

	i := 0
	j := len(copyToSort) - 1

	iv := copyToSort[i]
	jv := copyToSort[j]

	for iv <= jv {
		if iv+jv == sum {
			return true
		}
		if iv+jv < sum {
			iv++
			continue
		}
		if iv+jv > sum {
			jv--
			continue
		}
	}
	return false
}
