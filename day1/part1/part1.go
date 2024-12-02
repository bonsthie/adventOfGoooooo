package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(name string) string {
	data, err := os.ReadFile(name)
	check(err)
	return string(data)
}

func main() {
	buff := getFile("../data")

	splits := strings.Split(buff, "\n")

	var ls1 []int
	var ls2 []int
	for _, v := range splits {
		splits_num := strings.Split(v, "   ")

		if len(splits_num) == 1 {
			break
		}
		num, _ := strconv.Atoi(splits_num[0])
		ls1 = append(ls1, num)
		num, _ = strconv.Atoi(splits_num[1])
		ls2 = append(ls2, num)

	}
	sort.Ints(ls1)
	sort.Ints(ls2)

	var num_range int = 0
	for i, _ := range ls1 {
		num_range += int(math.Abs(float64(ls1[i] - ls2[i])))
	}
	fmt.Printf("%v\n", num_range)
}
