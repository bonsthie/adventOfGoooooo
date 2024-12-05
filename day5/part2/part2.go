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

func swap(a, b *int) {
	*a, *b = *b, *a
}

func getFile(name string) []byte {
	data, err := os.ReadFile(name)
	check(err)
	return data
}

func createOrderMap(matches [][][]byte) map[int][]int {
	MapOrder := make(map[int][]int)

	for _, v := range matches {
		fmt.Printf("v: %s\n", v)
		index, _ := strconv.Atoi(string(v[1]))
		num, _ := strconv.Atoi(string(v[2]))
		MapOrder[index] = append(MapOrder[index], num)
	}
	return MapOrder
}

func createList(strList string) []int {
	strsList := strings.Split(strList, ",")

	list := make([]int, len(strsList))
	for i, v := range strsList {
		list[i], _ = strconv.Atoi(v)
	}
	return list
}

func contain(lst1, lst2 []int) int {
	for i, v := range lst1 {
		for _, v2 := range lst2 {
			if v == v2 {
				return i
			}
		}
	}
	return -1
}

func validList(intList []int, orderMap map[int][]int) int {

	for i, v := range intList {
		if i == 0 {
			continue
		}
		if contain(intList[:i], orderMap[v]) != -1 {
			return 0
		}
	}
	return 1
}

func invalidList(intList []int, orderMap map[int][]int) int {
	for validList(intList, orderMap) == 0 {
		for i, v := range intList {
			offset := contain(intList[:i], orderMap[v])
			if offset != -1 {
				swap(&intList[i], &intList[offset])
			}
		}
	}
	fmt.Printf("intList: %v\n", intList)
	return intList[(len(intList) / 2)]
}

func main() {
	file := getFile("../data")

	regOrder := regexp.MustCompile(`(\d+)\|(\d+)`)

	matchesOrder := regOrder.FindAllSubmatch(file, -1)

	matchesList := strings.Split(string(file[strings.Index(string(file), "\n\n")+2:]), "\n")

	OrderMap := createOrderMap(matchesOrder)

	count := 0
	for _, v := range matchesList {
		if strings.TrimSpace(v) == "" {
			continue
		}
		intList := createList(v)
		if validList(intList, OrderMap) == 0 {
			fmt.Printf("invalid : %v\n", v)
			count += invalidList(intList, OrderMap)
		}
	}
	fmt.Printf("count: %v\n", count)
}
