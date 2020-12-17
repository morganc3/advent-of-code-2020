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
	f, _ := os.Open("day10_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]
	lines_ints := make([]int, len(lines))
	for i, e := range lines {
		int_val, _ := strconv.Atoi(e)
		lines_ints[i] = int_val
	}

	sort.Ints(lines_ints)
	lines_ints = append([]int{0}, lines_ints...) // prepend 0 (adapter joltage)

	myAdapter := lines_ints[len(lines_ints)-1] + 3
	lines_ints = append(lines_ints, myAdapter)
	onesCount := 0
	threesCount := 0

	for i := range lines_ints {
		// check next 3 elements (the only elements where a difference of 1 or 3 will be possible)

		foundNext := false
		j := 0
		for !foundNext && j < 3 {
			checkNext(lines_ints, i, i+j, &onesCount, &threesCount)
			j++
		}
	}

	fmt.Println(onesCount * threesCount)
}

func checkNext(nums []int, currIndex, targetIndex int, onesCount, threesCount *int) bool {
	if targetIndex >= len(nums) {
		return false
	}

	diff := nums[targetIndex] - nums[currIndex]
	switch diff {
	case 1:
		*onesCount++
		return true
	case 3:
		*threesCount++
		return true
	}
	return false
}
