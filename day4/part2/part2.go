package main

import (
	"fmt"
	"os"
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

func check_field(fields []string, i int, j int) int {
    // Ensure indices are within bounds
    if i-1 < 0 || i+1 >= len(fields) || j-1 < 0 || j+1 >= len(fields[0]) {
        return 0
    }

    seq1 := string(fields[i-1][j-1]) + string(fields[i][j]) + string(fields[i+1][j+1])
    seq2 := string(fields[i-1][j+1]) + string(fields[i][j]) + string(fields[i+1][j-1])

    validSequences := map[string]bool{"MAS": true, "SAM": true}

    validSeq1 := validSequences[seq1]
    validSeq2 := validSequences[seq2]

    if validSeq1 && validSeq2 {
        return 1
    }
    return 0
}

func main() {
	file := getFile("../data")
	fields := strings.Fields(string(file))

	height := len(fields)
	weight := len(fields[0])
	count := 0

	for i := 1; i < height-1; i++ {
		for j := 1; j < weight-1; j++ {
			if fields[i][j] != 'A' {
				continue
			}
			count += check_field(fields, i, j)

		}
	}
	fmt.Printf("count: %v\n", count)
}
