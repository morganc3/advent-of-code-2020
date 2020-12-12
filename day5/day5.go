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
	i := 0
	j := len(lines_ints) - 1
	for i <= j {
		if lines_ints[i]+lines_ints[j] == 2020 {
			fmt.Println(lines_ints[i] * lines_ints[j])
			return
		}
		if lines_ints[i]+lines_ints[j] > 2020 {
			j--
			continue
		}
		if lines_ints[i]+lines_ints[j] < 2020 {
			i++
			continue
		}
	}

}
