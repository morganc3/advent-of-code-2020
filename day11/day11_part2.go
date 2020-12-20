package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var spaces [][]string

const (
	EMPTY    = "L"
	OCCUPIED = "#"
	FLOOR    = "."
)

func main() {
	f, _ := os.Open("day11_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	spaces = make([][]string, len(lines))
	// parse lines into 2d array
	for i := range lines {
		spaces[i] = make([]string, len(lines[0]))
		for j := range lines[i] {
			spaces[i][j] = string(lines[i][j])
		}
	}

	foundEqual := false
	for !foundEqual {
		// updates spaces in place, previous value
		// from before round is performed is stored in prev
		prev := performRound(spaces)
		foundEqual = checkEqual(prev, spaces)
	}

	fmt.Println(countAllOccupied(spaces))

}

func printBoard(spaces [][]string) {
	for _, line := range spaces {
		fmt.Println(line)
	}
	fmt.Println("")
}

func countAllOccupied(spaces [][]string) int {
	count := 0
	for y := range spaces {
		for x := range spaces[y] {
			if spaces[y][x] == OCCUPIED {
				count++
			}
		}
	}
	return count
}

func checkEqual(seats1, seats2 [][]string) bool {
	for y := range seats1 {
		for x := range seats1[y] {
			if seats1[y][x] != seats2[y][x] {
				return false
			}
		}
	}
	return true
}

// returns previous value of spaces, updates
// parameter in place
func performRound(spaces [][]string) [][]string {
	prevCopy := createCopy(spaces)
	// process empty spacess, turning to occupied if necessary
	for y := range spaces {
		for x := range spaces[y] {
			if spaces[y][x] == EMPTY && processEmpty(y, x, prevCopy) {
				spaces[y][x] = OCCUPIED
			}
			if spaces[y][x] == OCCUPIED && processOccupied(y, x, prevCopy) {
				spaces[y][x] = EMPTY
			}
		}
	}
	return prevCopy
}

func createCopy(spaces [][]string) [][]string {
	prevCopy := make([][]string, len(spaces))
	for i := range spaces {
		prevCopy[i] = make([]string, len(spaces[i]))
		copy(prevCopy[i], spaces[i])
	}
	return prevCopy
}

// returns value of space at x,y coords, returns ""
// if either index is invalid
func getSpace(y, x int, spaces [][]string) string {
	if x < 0 || x >= len(spaces[0]) {
		return ""
	}

	if y < 0 || y >= len(spaces) {
		return ""
	}

	return spaces[y][x]

}

// return count of occupied seats that
// can be seen from the position
func occupiedSeatsSeen(y, x int, spaces [][]string) int {
	pos := FLOOR
	count := 0
	directions := []string{"up", "down", "left", "right", "upRight", "downRight", "upLeft", "downLeft"}

	for _, d := range directions {
		pos = func(_y, _x int, direction string) string {
			pos = FLOOR
			for pos == FLOOR { // continue until we get to something that isn't floor
				switch direction {
				case "up":
					_y--
					pos = getSpace(_y, _x, spaces)
				case "down":
					_y++
					pos = getSpace(_y, _x, spaces)
				case "left":
					_x--
					pos = getSpace(_y, _x, spaces)
				case "right":
					_x++
					pos = getSpace(_y, _x, spaces)
				case "upRight":
					_y--
					_x++
					pos = getSpace(_y, _x, spaces)
				case "downRight":
					_y++
					_x++
					pos = getSpace(_y, _x, spaces)
				case "upLeft":
					_y--
					_x--
					pos = getSpace(_y, _x, spaces)
				case "downLeft":
					_x--
					_y++
					pos = getSpace(_y, _x, spaces)
				}
				if pos == OCCUPIED {
					count++
				}
			}
			return pos
		}(y, x, d)
	}

	return count
}

// if a seat is empty, change it to occupied
// if all spaces around it are empty
// returns true if empty seat should be flipped
func processEmpty(y, x int, spaces [][]string) bool {
	if occupiedSeatsSeen(y, x, spaces) != 0 {
		return false
	}

	return true
}

// returns true if filled seat should be flipped
func processOccupied(y, x int, spaces [][]string) bool {
	if occupiedSeatsSeen(y, x, spaces) >= 5 {
		return true
	}
	return false
}
