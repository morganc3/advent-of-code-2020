package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Bag struct {
	Color    string
	Contains map[string]int
}

func main() {
	f, _ := os.Open("day7_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n")
	lines = lines[0 : len(lines)-1]
	bags := parseLines(lines)

	queue := []string{"shiny gold"}

	count := 0
	alreadyCounted := []string{}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		alreadyCounted = append(alreadyCounted, curr)
		currCount, bs := directBagCount(curr, bags)

		badCount := 0 // keep track of how many returned we've already counted
		for _, b := range bs {
			// make sure we haven't counted the bag returned already, and
			// that it's not in our queue
			if !contains(b, alreadyCounted) && !contains(b, queue) {
				queue = append(queue, b)
			} else {
				badCount++
			}
		}
		count += currCount - badCount
	}
	fmt.Println(count)
}

func contains(item string, list []string) bool {
	for _, e := range list {
		if e == item {
			return true
		}
	}
	return false
}

// gets bags that can directly hold the color
func directBagCount(color string, bags []Bag) (int, []string) {
	count := 0
	var directBags []string
	for _, bag := range bags {
		for k, v := range bag.Contains {
			if k == color && v > 0 {
				count++
				directBags = append(directBags, bag.Color)
			}
		}
	}
	return count, directBags
}

func parseLines(lines []string) []Bag {
	var bags []Bag
	for _, line := range lines {
		bags = append(bags, parseLine(line))
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
