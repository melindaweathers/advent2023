package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards       string
	bid         int
	mainType    rune
	sortableVal string
}

// 2 3 4 5 6 7 8 9 T J Q K A
// 2 3 4 5 6 7 8 9 A B C D E
func sortableHand(s string) string {
	var m = map[rune]rune{
		'T': 'A',
		'J': 'B',
		'Q': 'C',
		'K': 'D',
		'A': 'E',
	}

	rtn := ""
	var sortableVal rune
	for _, c := range s {
		v, ok := m[c]
		if ok {
			sortableVal = v
		} else {
			sortableVal = c
		}
		rtn = rtn + string(sortableVal)
	}
	return rtn
}

func handType(s string) rune {
	cardCounts := map[rune]int{}
	var t rune
	for _, card := range s {
		cardCounts[card]++
	}
	countsOnly := []int{}
	for _, value := range cardCounts {
		countsOnly = append(countsOnly, value)
	}
	sort.Ints(countsOnly)
	if len(countsOnly) == 1 {
		// 5 Five of a kind
		t = 'G'
	} else if len(countsOnly) == 4 {
		// 2, 1, 1, 1 One pair
		t = 'B'
	} else if len(countsOnly) == 5 {
		// 1, 1, 1, 1, 1 High card
		t = 'A'
	} else if reflect.DeepEqual(countsOnly, []int{1, 1, 3}) {
		// 3, 1, 1 Three of a kind
		t = 'D'
	} else if len(countsOnly) == 3 {
		// 2, 2, 1 Two pair
		t = 'C'
	} else if reflect.DeepEqual(countsOnly, []int{1, 4}) {
		// 4, 1 Four of a kind
		t = 'F'
	} else {
		// 3, 2 Full house
		t = 'E'
	}
	return t
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func readHands(filename string) []Hand {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	hands := []Hand{}
	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, " ")
		handType := handType(parts[0])
		hands = append(hands, Hand{
			cards:       parts[0],
			bid:         toInt(parts[1]),
			mainType:    handType,
			sortableVal: string(handType) + sortableHand(parts[0]),
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[j].sortableVal > hands[i].sortableVal
	})
	return hands
}

func totalWinnings(filename string) int {
	hands := readHands(filename)
	// fmt.Println(hands)
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}

func main() {
	fmt.Println(totalWinnings("./input-test.txt"))
	fmt.Println(totalWinnings("./input-test2.txt"))
	fmt.Println(totalWinnings("./input.txt"))
}
