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
	position[0] = 0
	position[1] = 0

	treeCount := 0

	for position[1] < len(lines)-1 {
		move(position)

		currChar := string(lines[position[1]][position[0]])
		if currChar == "#" {
			treeCount++
		}
	}
	fmt.Println(treeCount)
}

func move(pos []int) {
	pos[0] = (pos[0] + 3) % 31 // how many columns we've moved right
	pos[1] += 1                // how many rows we've moved down
}
