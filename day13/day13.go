package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func isReflection(lines []string, pos int) bool {
	i := pos - 1
	j := pos
	for i >= 0 && j < len(lines) {
		if lines[i] != lines[j] {
			return false
		}
		i--
		j++
	}
	return true
}

func findReflection(lines []string, skipVal int) int {
	for i := 1; i < len(lines); i++ {
		if i != skipVal && isReflection(lines, i) {
			return i
		}
	}
	return 0
}

func smudgedLines(lines []string, smudgedI int, smudgedJ int) []string {
	smudged := []string{}
	for i, line := range lines {
		if i == smudgedI {
			newChar := '.'
			if line[smudgedJ] == '.' {
				newChar = '#'
			}
			newLine := []rune(line)
			newLine[smudgedJ] = newChar
			smudged = append(smudged, string(newLine))
		} else {
			smudged = append(smudged, line)
		}
	}
	return smudged
}

func findSmudgeReflection(lines []string, oldReflection int) int {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			reflection := findReflection(smudgedLines(lines, i, j), oldReflection)
			if reflection > 0 {
				return reflection
			}
		}
	}
	return 0
}

func findReflections(filename string) (int, int) {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	reflections := 0
	smudgeReflections := 0
	for _, lineBlock := range strings.Split(string(b), "\n\n") {
		rows := strings.Split(lineBlock, "\n")
		cols := make([]string, len(rows[0]))
		for _, row := range rows {
			for c, char := range row {
				cols[c] = cols[c] + string(char)
			}
		}
		rowReflection := findReflection(rows, -1)
		colReflection := findReflection(cols, -1)
		reflection := 100*rowReflection + colReflection
		smudgeReflection := 100*findSmudgeReflection(rows, rowReflection) + findSmudgeReflection(cols, colReflection)
		reflections += reflection
		fmt.Println(smudgeReflection)
		smudgeReflections += smudgeReflection
	}

	return reflections, smudgeReflections
}

func main() {
	fmt.Println(findReflections("./input-test.txt"))
	fmt.Println(findReflections("./input-test2.txt"))
	fmt.Println(findReflections("./input.txt"))
}
