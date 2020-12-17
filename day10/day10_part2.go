package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var cachedAnswers map[int]int

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

	// create cache of answers so we don't repeat DFS work
	cachedAnswers = make(map[int]int)

	fmt.Println(dfs(lines_ints, 0))

}

// do dfs, saving the value for each number
// for instance if we branch left to 4, then save that for later. if tree attempts to branch to 4
// ANYWHERE else, use previously stored value

func dfs(nums []int, index int) int {
	// Dynamic programming, cache DFS answers
	if val, ok := cachedAnswers[index]; ok {
		return val
	}

	neighbors := getConnected(nums, index)
	neighborCount := len(neighbors)

	if neighborCount == 0 {
		return 1
	}

	return dfsHelper(nums, neighbors)

}

// call dfs on each neighbor, sum the results
func dfsHelper(nums []int, indices []int) int {
	var total int
	for _, ind := range indices {
		dfsResult := dfs(nums, ind)
		cachedAnswers[ind] = dfsResult
		total += dfsResult
	}
	return total
}

// get indices of "neighbors" that are connected

func getConnected(nums []int, index int) []int {
	currNum := nums[index]
	offset := 1
	ret := []int{}
	for index+offset < len(nums) && offset < 4 {
		diff := nums[index+offset] - currNum
		if diff >= 1 && diff <= 3 {
			ret = append(ret, index+offset)
		}
		offset++
	}
	return ret
}
