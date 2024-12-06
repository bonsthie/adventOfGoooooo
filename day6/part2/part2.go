package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	cursorType = "^>v<"
)

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

func mapSize(Map []string) (int, int) {
	height := len(Map)
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
		if y-1 < 0 {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}

		if Map[y-1][x] == '#' {
			Map[y] = replaceChar(Map[y], x, '>')
		} else {
			Map[y] = handleIntersection(Map[y], x)
			Map[y-1] = replaceChar(Map[y-1], x, '^')
		}
		return true
	}
}

func handleRight(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		if x+1 >= width {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}

		if Map[y][x+1] == '#' {
			Map[y] = replaceChar(Map[y], x, 'v')
		} else {
			Map[y] = handleIntersection(Map[y], x)
			Map[y] = replaceChar(Map[y], x+1, '>')
		}
		return true
	}
}

func handleDown(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		if y+1 >= height {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}

		if Map[y+1][x] == '#' {
			Map[y] = replaceChar(Map[y], x, '<')
		} else {
			Map[y] = handleIntersection(Map[y], x)
			Map[y+1] = replaceChar(Map[y+1], x, 'v')
		}
		return true
	}
}

func handleLeft(x, y int) func(Map []string, width, height int) bool {
	return func(Map []string, width, height int) bool {
		if x-1 < 0 {
			Map[y] = replaceChar(Map[y], x, 'X')
			return false
		}

		if Map[y][x-1] == '#' {
			Map[y] = replaceChar(Map[y], x, '^')
		} else {
			Map[y] = handleIntersection(Map[y], x)
			Map[y] = replaceChar(Map[y], x-1, '<')
		}
		return true
	}
}

func handleIntersection(line string, x int) string {
	current := line[x]
	if current == '|' || current == '-' || current == '+' {
		return replaceChar(line, x, '+')
	}
	return replaceChar(line, x, '|')
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

func countX(Map []string) int {
	count := 0
	for _, line := range Map {
		count += strings.Count(line, "X")
	}
	return count
}

func main() {
	// Read file and create map
	file := string(getFile("../data"))
	Map := strings.Split(file, "\n")

	width, height := mapSize(Map)

	for {
		cursorFunc, err := findCursorPos(Map)
		if err != nil {
			break
		}
		if !cursorFunc(Map, width, height) {
			break
		}
	}

	// Print final map and count 'X'
	for _, line := range Map {
		fmt.Println(line)
	}
	fmt.Printf("Number of X: %d\n", countX(Map))
}
