package main

import (
	"bufio"
	"errors"
	"fmt"
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

func getBothNumber(line string) (int, int, error) {
	fields := strings.Fields(line)

	if len(fields) < 2 {
		return 0, 0, errors.New("wrong format need 2 numbers\n")
	}

    num1, err := strconv.Atoi(fields[0])
	if err != nil {
		return 0, 0, err
	}
	num2, err := strconv.Atoi(fields[1])
	if err != nil {
		return 0, 0, err
	}
	return num1, num2, nil
}

func main() {
	fd := getFile("../data")
	defer fd.Close()

	var lst1 []int
	m_lst2 := make(map[int]int)
	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		num1, num2, err := getBothNumber(line)
		if (err != nil) {
			break
		}
		lst1 = append(lst1, num1)
		m_lst2[num2]++
	}
	
	var sum int
	for _, v := range(lst1) {
		sum += v * m_lst2[v]
	}
	fmt.Printf("sum: %v\n", sum)
}
