package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func shiftNorth(lines []string) ([]string, int) {
	numShifted := 0
	shiftedLines := []string{}
	shiftedLines = append(shiftedLines, lines[0])

	for i := 1; i < len(lines); i++ {
		prevLine := []rune(shiftedLines[i-1])
		curLine := []rune(lines[i])

		for j := 0; j < len(prevLine); j++ {
			if curLine[j] == 'O' && prevLine[j] == '.' {
				prevLine[j] = 'O'
				curLine[j] = '.'
				numShifted++
			}
		}

		shiftedLines[i-1] = string(prevLine)
		shiftedLines = append(shiftedLines, string(curLine))
	}

	return shiftedLines, numShifted
}

func countLoad(lines []string) int {
	load := 0
	for i, line := range lines {
		for _, char := range line {
			if char == 'O' {
				load += len(lines) - i
			}
		}
	}

	return load
}

func northLoad(filename string) int {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	numShifted := 1
	for numShifted > 0 {
		lines, numShifted = shiftNorth(lines)
	}

	return countLoad(lines)
}

func main() {
	fmt.Println(northLoad("./input-test.txt"))
	fmt.Println(northLoad("./input.txt"))
}
