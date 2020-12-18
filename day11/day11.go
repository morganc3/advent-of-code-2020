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
		prev := performRound(spaces)
		foundEqual = checkEqual(prev, spaces)
	}

	fmt.Println(countAllOccupied(spaces))

	// for _, line := range spaces {
	// 	fmt.Println(line)
	// }

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

// returns previous value of spaces
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

func getSpace(y, x int, spaces [][]string) string {
	if x < 0 || x >= len(spaces[0]) {
		return ""
	}

	if y < 0 || y >= len(spaces) {
		return ""
	}

	return spaces[y][x]

}

// if a seat is empty, change it to occupied
// if all spaces around it are empty
// returns true if empty seat should be flipped
func processEmpty(y, x int, spaces [][]string) bool {

	up := getSpace(y-1, x, spaces) == OCCUPIED
	down := getSpace(y+1, x, spaces) == OCCUPIED
	right := getSpace(y, x+1, spaces) == OCCUPIED
	left := getSpace(y, x-1, spaces) == OCCUPIED

	if up || down || right || left {
		return false // do not change seat to occupied
	}

	upRight := getSpace(y-1, x+1, spaces) == OCCUPIED
	downRight := getSpace(y+1, x+1, spaces) == OCCUPIED
	upLeft := getSpace(y-1, x-1, spaces) == OCCUPIED
	downLeft := getSpace(y+1, x-1, spaces) == OCCUPIED

	if upRight || downRight || upLeft || downLeft {
		return false // do not change seat to occupied
	}

	return true
}

// returns true if filled seat should be flipped
func processOccupied(y, x int, spaces [][]string) bool {
	up := getSpace(y-1, x, spaces) == OCCUPIED
	down := getSpace(y+1, x, spaces) == OCCUPIED
	right := getSpace(y, x+1, spaces) == OCCUPIED
	left := getSpace(y, x-1, spaces) == OCCUPIED
	upRight := getSpace(y-1, x+1, spaces) == OCCUPIED
	downRight := getSpace(y+1, x+1, spaces) == OCCUPIED
	upLeft := getSpace(y-1, x-1, spaces) == OCCUPIED
	downLeft := getSpace(y+1, x-1, spaces) == OCCUPIED

	occupiedSeats := []bool{up, down, right, left, upRight, downRight, upLeft, downLeft}
	occupiedCount := 0
	for _, seat := range occupiedSeats {
		if seat {
			occupiedCount++
		}
	}
	if occupiedCount >= 4 {
		return true
	}

	return false
}
