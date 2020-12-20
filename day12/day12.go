package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	NORTH   = "N"
	SOUTH   = "S"
	EAST    = "E"
	WEST    = "W"
	FORWARD = "F"
	RIGHT   = "R"
	LEFT    = "L"
)

// starts facing east
type Ship struct {
	Direction string
	East      int
	North     int
}

var directions []string = []string{EAST, SOUTH, WEST, NORTH}

func main() {
	f, _ := os.Open("day12_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	ship := Ship{Direction: EAST, East: 0, North: 0}
	for _, line := range lines {
		ship.parseAndMove(line)
	}

	if ship.East < 0 {
		ship.East *= -1
	}
	if ship.North < 0 {
		ship.North *= -1
	}
	fmt.Println(ship.East + ship.North)

}

func (s *Ship) parseAndMove(line string) {
	action := string(line[0])
	valueStr := line[1:]
	value, _ := strconv.Atoi(valueStr)

	switch action {
	case RIGHT:
		s.turn(RIGHT, value)
	case LEFT:
		s.turn(LEFT, value)
	case FORWARD, NORTH, SOUTH, EAST, WEST:
		s.move(action, value)
	}
}

func (s *Ship) move(direction string, value int) {
	switch direction {
	case FORWARD:
		s.move(s.Direction, value)
	case NORTH:
		s.North += value
	case SOUTH:
		s.North -= value
	case EAST:
		s.East += value
	case WEST:
		s.East -= value
	}
}

func getDirectionsIndex(direction string) int {
	for i, d := range directions {
		if d == direction {
			return i
		}
	}
	return -1
}

func (s *Ship) turn(direction string, degrees int) {
	rotations := degrees / 90
	currDirection := s.Direction
	currDirectionIndex := getDirectionsIndex(currDirection)
	var newDirection string

	for rotations > 0 {

		if direction == RIGHT {
			currDirectionIndex++ // rotate by 90 degrees
		} else {
			currDirectionIndex--
		}

		if currDirectionIndex >= len(directions) {
			currDirectionIndex = 0
		}
		if currDirectionIndex < 0 {
			currDirectionIndex = len(directions) - 1
		}
		newDirection = directions[currDirectionIndex]
		rotations--
	}
	s.Direction = newDirection
}
