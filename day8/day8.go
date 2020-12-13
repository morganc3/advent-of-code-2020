package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Name     string
	Executed bool
	Value    int
}

func main() {
	f, _ := os.Open("day8_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]

	instructions := parseInstructions(lines)
	answer := runCode(instructions)
	fmt.Println(answer)
}

func parseInstructions(lines []string) []Instruction {
	ret := []Instruction{}
	for _, line := range lines {
		contents := strings.Split(line, " ")
		instName := contents[0]
		instVal, _ := strconv.Atoi(string(contents[1][1:]))

		negative := (string(contents[1][0]) == "-")
		if negative {
			instVal = -1 * instVal
		}

		instruction := Instruction{Name: instName, Executed: false, Value: instVal}
		ret = append(ret, instruction)
	}

	return ret
}

func runCode(instructions []Instruction) int {
	instructionCounter := 0
	accumulator := 0

	for {
		instructionCounterPrev := instructionCounter
		currInstruction := instructions[instructionCounter]

		if currInstruction.Executed {
			return accumulator
		}
		switch currInstruction.Name {
		case "jmp":
			instructionCounter += currInstruction.Value
		case "nop":
			instructionCounter++
		case "acc":
			accumulator += currInstruction.Value
			instructionCounter++
		}
		instructions[instructionCounterPrev].Executed = true
	}
}
