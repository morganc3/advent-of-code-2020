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

	count := 0
	for _, g := range groups {
		count += countGroup(g)
	}
	fmt.Println(count)
}

func countGroup(group string) int {
	answers := make(map[rune]bool)
	for _, c := range group {
		matched, _ := regexp.MatchString(`^[a-z]$`, string(c))
		if matched {
			answers[c] = true
		}
	}
	return len(answers)
}
