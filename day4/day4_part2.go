package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("day4_input.txt")
	b, _ := ioutil.ReadAll(f)
	input_string := string(b)
	lines := strings.Split(input_string, "\n\n")

	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}

	count := 0
	for _, line := range lines {
		if validateFields(line, fields) {
			count++
		}
	}
	fmt.Println(count)
}

func getFieldValue(line, key string) string {
	fields := strings.Fields(line)
	for _, f := range fields {
		if strings.Contains(f, key+":") {
			kv := strings.Split(f, ":")
			val := kv[1]
			return val
		}
	}
	return ""
}

func fieldExists(line, key string) bool {
	return strings.Contains(line, fmt.Sprintf("%s:", key))
}

func validRange(val, lower, upper int) bool {
	if val < lower || val > upper {
		return false
	}
	return true
}

func isValid(line, field string) bool {
	v := getFieldValue(line, field)

	switch field {
	case "byr":
		val, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return validRange(val, 1920, 2002)
	case "iyr":
		val, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return validRange(val, 2010, 2020)
	case "eyr":
		val, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		return validRange(val, 2020, 2030)
	case "hgt":
		cm := strings.Contains(v, "cm") //check if "cm" or "in"
		v = strings.ReplaceAll(v, "cm", "")
		v = strings.ReplaceAll(v, "in", "")

		val, err := strconv.Atoi(v)
		if err != nil {
			return false
		}
		if cm {
			return validRange(val, 150, 193)
		}
		return validRange(val, 59, 76)
	case "hcl":
		if len(v) != 7 {
			return false
		}
		v = v[1:] // remove preceeding "#"
		_, err := hex.DecodeString(v)
		if err != nil {
			return false // contains a non hex character
		}
		return true
	case "ecl":
		eyeColors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, c := range eyeColors {
			if v == c {
				return true
			}
		}
		return false
	case "pid":
		_, err := strconv.Atoi(v)
		if err != nil || len(v) != 9 {
			return false
		}
		return true
	case "cid":
		return true

	}
	return false
}

func validateFields(line string, fields []string) bool {
	for _, field := range fields {
		if !fieldExists(line, field) && field != "cid" {
			return false
		}
		if !isValid(line, field) {
			return false
		}
	}
	return true
}
