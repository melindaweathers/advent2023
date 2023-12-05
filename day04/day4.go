package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func process(filename string, processor func(string, map[int]int) int) int {
	sum := 0
	state := map[int]int{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += processor(scanner.Text(), state)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return sum
}

// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
func matchingCards(line string) int {
	cards := 0
	outer := strings.Split(line, ": ")
	inner := strings.Split(outer[1], " | ")

	winners := map[string]bool{}
	for _, winner := range strings.Fields(inner[0]) {
		winners[winner] = true
	}

	for _, ourNumber := range strings.Fields(inner[1]) {
		if winners[ourNumber] {
			cards += 1
		}
	}
	return cards
}

func cardScore(line string, _ map[int]int) int {
	cards := matchingCards(line)
	if cards == 0 {
		return 0
	}
	score := 1
	for i := 2; i <= cards; i++ {
		score *= 2
	}
	return score
}

func sumCards(line string, state map[int]int) int {
	outer := strings.Split(line, ": ")
	cardNum, _ := strconv.Atoi(strings.Fields(outer[0])[1])
	cards := matchingCards(line)

	for i := 1; i <= cards; i++ {
		state[cardNum+i] = state[cardNum+i] + state[cardNum] + 1
	}

	return 1 + state[cardNum]
}

func main() {
	fmt.Println(process("./input-test.txt", cardScore))
	fmt.Println(process("./input.txt", cardScore))
	fmt.Println(process("./input-test.txt", sumCards))
	fmt.Println(process("./input.txt", sumCards))
}
