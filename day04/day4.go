package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
func sumForLine(line string) int {
	score := 0
	outer := strings.Split(line, ": ")
	inner := strings.Split(outer[1], " | ")

	winners := map[string]bool{}
	for _, winner := range strings.Fields(inner[0]) {
		winners[winner] = true
	}

	for _, ourNumber := range strings.Fields(inner[1]) {
		if winners[ourNumber] {
			if score == 0 {
				score = 1
			} else {
				score = score * 2
			}
		}
	}
	return score
}

func main() {
	fmt.Println(process("./input-test.txt", sumForLine))
	fmt.Println(process("./input.txt", sumForLine))
}
