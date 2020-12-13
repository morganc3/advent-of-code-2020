package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// bag we parsed from a line of input text
type Bag struct {
	Color    string
	Contains map[string]int
}

// node of tree structure
type BagNode struct {
	Color string
	Nodes []*BagNode
}

var BagCount int = 0
var Bags map[string]Bag

func main() {
	f, _ := os.Open("day7_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]
	Bags = parseLines(lines)

	Tree := createTree("shiny gold")
	_ = Tree

	fmt.Println(BagCount)
}

func createTree(color string) *BagNode {
	currBag := Bags[color]
	if len(currBag.Contains) == 0 {
		return nil
	}

	currBagNode := BagNode{Color: color}
	children := []*BagNode{}

	for k, v := range currBag.Contains {
		for i := 0; i < v; i++ {
			children = append(children, createTree(k))
			BagCount++
		}
	}
	currBagNode.Nodes = children
	return &currBagNode
}

func parseLines(lines []string) map[string]Bag {
	bags := make(map[string]Bag)
	for _, line := range lines {
		b := parseLine(line)
		bags[b.Color] = b
	}
	return bags
}

func parseLine(line string) Bag {
	words := strings.Split(line, " ")
	contains := make(map[string]int)
	bag := Bag{
		Color:    fmt.Sprintf("%s %s", words[0], words[1]), // e.g. light + green
		Contains: contains,
	}
	containedBagsString := words[4:]
	if containedBagsString[0] == "no" { // "contain no other bags"
		return bag
	}
	for i := 0; i < len(containedBagsString); i += 4 {
		bagColor := fmt.Sprintf("%s %s", containedBagsString[i+1], containedBagsString[i+2])
		bagCount, err := strconv.Atoi(containedBagsString[i])
		if err != nil {
			log.Fatal("failed to convert str to int, parsing error")
		}
		bag.Contains[bagColor] = bagCount
	}

	return bag
}
