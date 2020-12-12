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
	f, _ := os.Open("day1_input.txt")
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

	// could delete any numbers > 2020

	for i, pivot := range lines_ints {
		j := len(lines_ints) - 1
		k := 0
		for k < i && j > i {
			sum := pivot + lines_ints[k] + lines_ints[j]
			if sum == 2020 {
				fmt.Println(pivot, lines_ints[k], lines_ints[j])
				fmt.Println(pivot * lines_ints[k] * lines_ints[j])
				return
			}

			if sum > 2020 {
				j--
				continue
			}

			if sum < 2020 {
				k++
				continue
			}
		}
	}

}
