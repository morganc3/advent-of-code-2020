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

type Waypoint struct {
	Quadrant int
	East     int
	North    int
	WPShip   *Ship
}

// starts facing east
type Ship struct {
	East  int
	North int
}

func main() {
	f, _ := os.Open("day12_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	ship := Ship{East: 0, North: 0}
	waypoint := Waypoint{East: 10, North: 1, Quadrant: 1, WPShip: &ship}
	for _, line := range lines {
		waypoint.parseAndMove(line)
	}

	if ship.East < 0 {
		ship.East *= -1
	}
	if ship.North < 0 {
		ship.North *= -1
	}
	fmt.Println(ship.East + ship.North)
}

func (s *Ship) printShip() {
	fmt.Printf("Ship is at EAST: %d, NORTH: %d\n", s.East, s.North)
}

func (w *Waypoint) printWaypoint() {
	fmt.Printf("Waypoint is at EAST: %d, NORTH: %d (Quadrant %d)\n", w.East, w.North, w.Quadrant)
}

func (w *Waypoint) parseAndMove(line string) {
	action := string(line[0])
	valueStr := line[1:]
	value, _ := strconv.Atoi(valueStr)

	switch action {
	case RIGHT:
		w.turn(RIGHT, value)
	case LEFT:
		w.turn(LEFT, value)
	case FORWARD:
		w.moveShip(value)
	case NORTH, SOUTH, EAST, WEST:
		w.moveWaypoint(action, value)
	}
}

func (w *Waypoint) moveShip(value int) {
	ship := w.WPShip
	east := w.East
	north := w.North
	eastMovement := east * value
	northMovement := north * value

	ship.East += eastMovement
	ship.North += northMovement
}

func (w *Waypoint) moveWaypoint(direction string, value int) {
	switch direction {
	case NORTH:
		w.North += value
	case SOUTH:
		w.North -= value
	case EAST:
		w.East += value
	case WEST:
		w.East -= value
	}

	// update quadrant if need be
	if w.East > 0 && w.North > 0 {
		w.Quadrant = 1
	}
	if w.East > 0 && w.North < 0 {
		w.Quadrant = 4
	}
	if w.East < 0 && w.North > 0 {
		w.Quadrant = 2
	}
	if w.East < 0 && w.North < 0 {
		w.Quadrant = 3
	}
}

func (w *Waypoint) turn(direction string, degrees int) {
	for degrees > 0 {
		w.turn90(direction)
		degrees -= 90
	}
}

func (w *Waypoint) turn90(direction string) {
	q := w.Quadrant
	if direction == RIGHT {
		q--
	} else {
		q++
	}
	if q > 4 {
		q = 1
	}
	if q < 1 {
		q = 4
	}

	// swap
	tmp := w.East
	w.East = w.North
	w.North = tmp

	w.Quadrant = q

	if q == 1 {
		w.East = makePos(w.East)
		w.North = makePos(w.North)
	}
	if q == 2 {
		w.East = makeNeg(w.East)
		w.North = makePos(w.North)
	}
	if q == 3 {
		w.East = makeNeg(w.East)
		w.North = makeNeg(w.North)
	}
	if q == 4 {
		w.East = makePos(w.East)
		w.North = makeNeg(w.North)
	}
}

func makePos(in int) int {
	if in < 0 {
		in *= -1
	}
	return in
}

func makeNeg(in int) int {
	if in > 0 {
		in *= -1
	}
	return in
}
