package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cursorType = "^>v<"
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

func mapSize(Map []string) (int, int) {
	height := len(Map) - 1
	if height == 0 {
		return 0, 0
	}
	return len(Map[0]), height
}

func replaceChar(line string, x int, char rune) string {
	runeLine := []rune(line)
	runeLine[x] = char
	return string(runeLine)
}

func handleUp(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		fmt.Printf("Cursor is facing up (^) at (%d, %d) in a map of size (%d, %d)\n", x, y, width, height)
		if y-1 < 0 {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}
		if Map[y-1][x] == '#' {
			Map[y] = replaceChar(Map[y], x, '>')
		} else {
			Map[y] = replaceChar(Map[y], x, 'X')
			Map[y-1] = replaceChar(Map[y-1], x, '^')
		}
		return true
	}
}

func handleRight(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		fmt.Printf("Cursor is facing right (>) at (%d, %d) in a map of size (%d, %d)\n", x, y, width, height)
		if x+1 >= width {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}
		if Map[y][x+1] == '#' {
			Map[y] = replaceChar(Map[y], x, 'v')
		} else {
			Map[y] = replaceChar(Map[y], x, 'X')
			Map[y] = replaceChar(Map[y], x+1, '>')
		}
		return true
	}
}

func handleDown(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		fmt.Printf("Cursor is facing down (v) at (%d, %d) in a map of size (%d, %d)\n", x, y, width, height)
		if y+1 >= height {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}
		if Map[y+1][x] == '#' {
			Map[y] = replaceChar(Map[y], x, '<')
		} else {
			Map[y] = replaceChar(Map[y], x, 'X')
			Map[y+1] = replaceChar(Map[y+1], x, 'v')
		}
		return true
	}
}

func handleLeft(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		fmt.Printf("Cursor is facing left (<) at (%d, %d) in a map of size (%d, %d)\n", x, y, width, height)
		if x-1 < 0 {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}
		if Map[y][x-1] == '#' {
			Map[y] = replaceChar(Map[y], x, '^')
		} else {
			Map[y] = replaceChar(Map[y], x, 'X')
			Map[y] = replaceChar(Map[y], x-1, '<')
		}
		return true
	}
}

func findCursorPos(Map []string) (func(Map []string, width, height int) bool, error) {
	cursorFunctions := map[rune]func(x, y int) func(Map []string, width, height int) bool{
		'^': handleUp,
		'>': handleRight,
		'v': handleDown,
		'<': handleLeft,
	}

	for y, line := range Map {
		for _, c := range cursorType {
			x := strings.IndexRune(line, c)
			if x != -1 {
				return cursorFunctions[c](x, y), nil
			}
		}
	}
	return nil, fmt.Errorf("no cursor found")
}

func numberofX(Map []string) int {
	count := 0
	for _, v := range Map {
		count += strings.Count(v, "X")
	}
	return count
}

func main() {
	file := string(getFile("../data"))
	Map := strings.Split(file, "\n")

	width, height := mapSize(Map)

	for {
		cursorFunc, err := findCursorPos(Map)
		if err != nil {
			fmt.Printf("numberofX(Map): %v\n", numberofX(Map))
			return
		}
		if cursorFunc(Map, width, height) == false {
			fmt.Printf("numberofX(Map): %v\n", numberofX(Map))
			return
		}

	}
}
