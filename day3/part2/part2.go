package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func getFile(name string) []byte {
	data, err := os.ReadFile(name)
	check(err)
	return data
}

func main() {

	file := getFile("../data")
	reg := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don\'t\(\)|do\(\)`)

	matches := reg.FindAllSubmatch(file, -1)

	total := 0
	takecount := 1
	for i, match := range matches {
		fmt.Printf("Match %d: %s\n", i+1, match[0])
		if strings.Compare(string(match[0]), "don't()") == 0 {
			takecount = 0
			continue
		}
		if strings.Compare(string(match[0]), "do()") == 0 {
			takecount = 1
			continue
		}

		if takecount == 0 {
			continue
		}

		fmt.Printf("  Group 1: %s\n", match[1])
		fmt.Printf("  Group 2: %s\n", match[2])

		num1, _ := strconv.Atoi(string(match[1]))
		num2, _ := strconv.Atoi(string(match[2]))

		total += (num1 * num2)
	}
	fmt.Printf("total: %v\n", total)
}
