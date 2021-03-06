package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("day3_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	position := make([]int, 2)

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	treeProduct := 1
	for _, s := range slopes {
		position[0] = 0
		position[1] = 0
		treeCount := 0
		for position[1] < len(lines)-1 {
			move(position, s[0], s[1])

			currChar := string(lines[position[1]][position[0]])
			if currChar == "#" {
				treeCount++
			}
		}

		treeProduct *= treeCount
	}

	fmt.Println(treeProduct)

}

func move(pos []int, x, y int) {
	pos[0] = (pos[0] + x) % 31 // how many columns we've moved right
	pos[1] += y                // how many rows we've moved down
}
