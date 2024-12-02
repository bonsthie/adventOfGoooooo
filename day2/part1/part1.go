package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type CompareFunc func(a, b int) bool

func greaterThan(a, b int) bool {
	return a > b
}

func lessThan(a, b int) bool {
	return a < b
}

func getCmpFunc(fields []string) CompareFunc {
	a, err1 := strconv.Atoi(fields[0])
	b, err2 := strconv.Atoi(fields[1])

	if err1 != nil || err2 != nil {
		return nil
	}

	if a < b {
		return lessThan
	} else {
		return greaterThan
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(name string) *os.File {
	fd, err := os.Open(name)
	check(err)
	return fd
}

func isSafe(str string) uint {
	fields := strings.Fields(str)

	if len(fields) < 2 {
		return 1
	}

	cmp := getCmpFunc(fields)
	last, _ := strconv.Atoi(fields[0])
	for i, v := range fields {
		if i == 0 {
			continue
		}

		new, _ := strconv.Atoi(v)
		if cmp(last, new) == false {
			return 0
		}

		if int(math.Abs(float64(last - new))) > 3 {
			return 0
		}
		last = new
	}
	return 1
}

func main() {
	fd := getFile("../data")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	var safeCounter uint = 0
	for scanner.Scan() {
		line := scanner.Text()
		safeCounter += isSafe(line)
	}
	fmt.Printf("safeCounter: %v\n", safeCounter)
}
