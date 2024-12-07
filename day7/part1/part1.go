package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Calcul struct {
	Res int
	Num []int
}

// overkill but i have les crampte of coding this
func ParseCalculLine(line string) (*Calcul, error) {
	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid format: missing ':' separator")
	}

	res, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return nil, fmt.Errorf("invalid Res value '%s': %v", parts[0], err)
	}

	numStrings := strings.Fields(strings.TrimSpace(parts[1]))
	nums := make([]int, len(numStrings))
	for i, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s': %v", numStr, err)
		}
		nums[i] = num
	}

	return &Calcul{
		Res: res,
		Num: nums,
	}, nil
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getFile(name string) []byte {
	data, err := os.ReadFile(name)
	check(err)
	return data
}

func isValidCalc(calc *Calcul) bool {
	if (len(calc.Num) == 1) {
		
	}
}


func main() {
	file := string(getFile("../data"))

	lines := strings.Split(file, "\n")

	var count uint64 = 0
	for _, line := range lines {
		calc, _ := ParseCalculLine(line)
		if isValidCalc(calc) {
			count += uint64(calc.Res)
		}
		
	}
	fmt.Printf("count: %v\n", count)
}
