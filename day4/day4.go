package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("day4_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n\n")
	lines = lines[0:len(lines)]
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

	count := 0
	for _, line := range lines {
		if isValid(line, fields) {
			count++
		}
	}
	fmt.Println(count)

}

func fieldExists(line, key string) bool {
	return strings.Contains(line, fmt.Sprintf("%s:", key))
}

func isValid(line string, fields []string) bool {
	for _, f := range fields {
		fe := fieldExists(line, f)
		if !fe && f != "cid" {
			return false
		}
	}
	return true
}
