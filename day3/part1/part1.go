package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	reg := regexp.MustCompile(`/mul\((\d{1,3}),(\d{1,3})\)|don\'t\(\)|do\(\)/gm`)

	matches := reg.FindAllSubmatch(file, -1)

	total := 0
	for i, match := range matches {
		fmt.Printf("Match %d: %s\n", i+1, match[0]) 
		fmt.Printf("  Group 1: %s\n", match[1])
		fmt.Printf("  Group 2: %s\n", match[2])
		num1, _ := strconv.Atoi(string(match[1]))
		num2, _ := strconv.Atoi(string(match[2]))

		total += (num1 * num2)
	}
	fmt.Printf("total: %v\n", total)
}
