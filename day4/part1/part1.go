package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	word     = "XMAS"
	rev_word = "SAMX"
	word_len = len(word)
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

func isXMAS(str string) int {
	if str[:word_len] == word || str[:word_len] == rev_word {
		return 1
	}
	// fmt.Print("no\n")
	return 0
}

func vertical_check(fields []string, offset int) int {

	var str string
	for i := 0; i < word_len; i++ {
		str += string(fields[i][offset])
	}
	return (isXMAS(str))
}

func diag_right(fields []string, offset int) int {
	var str string
	for i := 0; i < word_len; i++ {
		str += string(fields[i][offset+i])
	}
	return (isXMAS(str))
}

func diag_left(fields []string, offset int) int {
	var str string
	for i := 0; i < word_len; i++ {
		str += string(fields[i][offset-i])
	}
	return (isXMAS(str))
}

func main() {
	file := getFile("../data")
	fields := strings.Fields(string(file))

	height := len(fields)
	weight := len(fields[0])
	count := 0

	for i := 0; i < height; i++ {
		for j := 0; j < weight; j++ {
			if fields[i][j] != 'X' && fields[i][j] != 'S' {
				continue
			}

			if j+word_len <= weight {

				count += isXMAS(fields[i][j:])

				if i+word_len <= height {
					count += diag_right(fields[i:], j)
				}

			}
			if i+word_len <= height {
				count += vertical_check(fields[i:], j)

				if j - word_len + 1>= 0 {
					count += diag_left(fields[i:], j)

				}
			}

		}
	}
	fmt.Printf("count: %v\n", count)
}

//s..s..s
//.a.a.a.
//..mmm..
//samxmas
//..mmm..
//.a.a.a.
//s..s..s
