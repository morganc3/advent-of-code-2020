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

	BADNUM := 1492208709

	for i := range lines_ints {
		lower, upper := checkSum(lines_ints, i, BADNUM)
		if lower != -1 {
			contiguousNumsForSum := lines_ints[lower : upper+1]
			sort.Ints(contiguousNumsForSum)
			fmt.Println(contiguousNumsForSum[0] + contiguousNumsForSum[len(contiguousNumsForSum)-1])
			return
		}
	}

}

// return lower, upper
func checkSum(nums []int, lower, sum int) (int, int) {
	currSum := 0
	currLower := lower
	for currLower < len(nums) {
		currSum += nums[currLower]
		if currSum == sum {
			return lower, currLower
		}
		if currSum > sum {
			return -1, -1
		}
		currLower++
	}
	return -1, -1
}
