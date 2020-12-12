package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, _ := os.Open("day6_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	groups := strings.Split(input_string, "\n\n")

	// remove trailing newline from last group
	// since we're going to split on newline later
	last := groups[len(groups)-1]
	last = last[:len(last)-1]
	groups[len(groups)-1] = last

	count := 0
	for _, g := range groups {
		count += countGroup(g)
	}
	fmt.Println(count)
}

func countGroup(group string) int {
	people := strings.Split(group, "\n")
	peopleCount := len(people)

	answers := make(map[rune]int)

	for _, p := range people {
		for _, c := range p {
			matched, _ := regexp.MatchString(`^[a-z]$`, string(c))
			if matched {
				answers[c]++
			}
		}
	}

	everyoneIsYesCount := 0
	for _, v := range answers {
		if v == peopleCount {
			everyoneIsYesCount++
		}
	}

	return everyoneIsYesCount
}
