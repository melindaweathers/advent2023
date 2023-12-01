package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process(filename string, processor func(string) int) int {
	sum := 0
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += processor(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

func sumForLine(line string) int {
	var first rune
	var last rune
	var sum int
	for _, letter := range line {
		val := int(letter - '0')
		if val >= 0 && val <= 9 {
			last = letter
			if first == 0 {
				first = letter
			}
		}
	}
	sum, _ = strconv.Atoi(string(first) + string(last))
	return sum
}

func numFromText(line string, i int) int {
	digits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	for k, v := range digits {
		if strings.HasPrefix(line[i:], k) {
			return v
		}
	}
	return -1
}

func sumForLinePartTwo(line string) int {
	var first rune
	var last rune
	var sum int
	var val int
	for i, letter := range line {
		val = int(letter - '0')
		if val < 0 || val > 9 {
			val = numFromText(line, i)
		}
		if val >= 0 && val <= 9 {
			last = rune(val) + '0'
			if first == 0 {
				first = last
			}
		}
	}
	sum, _ = strconv.Atoi(string(first) + string(last))
	return sum
}

func main() {
	fmt.Println(process("./input-test.txt", sumForLine))
	fmt.Println(process("./input.txt", sumForLine))
	fmt.Println(process("./input-test2.txt", sumForLinePartTwo))
	fmt.Println(process("./input.txt", sumForLinePartTwo))
}
