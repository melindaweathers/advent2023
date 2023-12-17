package main

import (
	"fmt"
	"log"
	"os"
	"slices"
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

func rotate(lines []string) []string {
	rotated := make([]string, len(lines[0]))
	for i := len(lines) - 1; i >= 0; i-- {
		for j := 0; j < len(lines[0]); j++ {
			rotated[j] = rotated[j] + string(lines[i][j])
		}
	}

	return rotated
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

func fullyShiftNorth(lines []string) ([]string, int) {
	numShifted := 1
	total := 0
	for numShifted > 0 {
		lines, numShifted = shiftNorth(lines)
		total += numShifted
	}
	return lines, total
}

func northLoad(filename string) int {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	lines, _ = fullyShiftNorth(lines)
	return countLoad(lines)
}

func fullSpin(filename string, numTimes int) int {
	b, err := os.ReadFile(filename)
	history := []string{}
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	loopAt := -1
	singleString := ""
	for times := 0; times < numTimes && loopAt < 0; times++ {
		for dir := 0; dir < 4; dir++ {
			lines, _ = fullyShiftNorth(lines)
			lines = rotate(lines)
		}
		singleString = strings.Join(lines, "~")
		loopAt = slices.Index(history, singleString)

		// fmt.Println(countLoad(lines))
		if loopAt < 0 {
			history = append(history, singleString)
		}
	}

	loopSize := len(history) - loopAt
	index := loopAt + ((numTimes - 1 - loopAt) % loopSize)

	return countLoad(strings.Split(history[index], "~"))
}

func main() {
	fmt.Println(northLoad("./input-test.txt"))
	fmt.Println(northLoad("./input.txt"))
	fmt.Println(fullSpin("./input-test.txt", 1000000000))
	fmt.Println(fullSpin("./input.txt", 1000000000))
}
