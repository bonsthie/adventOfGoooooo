package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

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

type CompareFunc func(a, b int) bool

func greaterThan(a, b int) bool {
	return a > b
}

func lessThan(a, b int) bool {
	return a < b
}

func getCmpFunc(a int, b int) CompareFunc {
	if a < b {
		return lessThan
	} else {
		return greaterThan
	}
}

func testNb(nb1 int, nb2 int, cmp CompareFunc) bool {
	diff := int(math.Abs(float64(nb1 - nb2)))
	return cmp(nb1, nb2) && diff <= 3
}

func isSort(tab []int) bool {
	if len(tab) < 2 {
		return true
	}
	cmp := getCmpFunc(tab[0], tab[1])
	for i := 1; i < len(tab); i++ {
		if !testNb(tab[i-1], tab[i], cmp) {
			return false
		}
	}
	return true
}

func isSafe(str string) uint {
	fields := strings.Fields(str)
	if len(fields) < 2 {
		return 1
	}

	tab := make([]int, len(fields))
	for i, v := range fields {
		tab[i], _ = strconv.Atoi(v)
	}

	for i, _ := range tab {
		tmp := make([]int, len(tab))
		copy(tmp, tab)
		tmp = append(tmp[:i], tmp[i+1:]...)
		if isSort(tmp) {
			return 1
		}
	}
	return 0
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
