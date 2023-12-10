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
		'O': '1',
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
	numJokers := 0
	for _, card := range s {
		cardCounts[card]++
	}
	countsOnly := []int{}
	for key, value := range cardCounts {
		if key == 'O' {
			numJokers = value
		} else {
			countsOnly = append(countsOnly, value)
		}
	}
	sort.Ints(countsOnly)
	if len(countsOnly) == 0 {
		countsOnly = []int{5}
	} else {
		countsOnly[len(countsOnly)-1] += numJokers
	}
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

func readHands(filename string, withJoker bool) []Hand {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	hands := []Hand{}
	for _, line := range strings.Split(string(b), "\n") {
		parts := strings.Split(line, " ")
		cards := ""
		if withJoker {
			for _, r := range parts[0] {
				if r == 'J' {
					cards = cards + "O"
				} else {
					cards = cards + string(r)
				}
			}
		} else {
			cards = parts[0]
		}
		handType := handType(cards)
		hands = append(hands, Hand{
			cards:       cards,
			bid:         toInt(parts[1]),
			mainType:    handType,
			sortableVal: string(handType) + sortableHand(cards),
		})
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[j].sortableVal > hands[i].sortableVal
	})
	return hands
}

func totalWinnings(filename string, withJoker bool) int {
	hands := readHands(filename, withJoker)
	total := 0
	for i, hand := range hands {
		total += (i + 1) * hand.bid
	}
	return total
}

func totalWinningsNoJoker(filename string) int {
	return totalWinnings(filename, false)
}

func totalWinningsWithJoker(filename string) int {
	return totalWinnings(filename, true)
}

func main() {
	fmt.Println(totalWinningsNoJoker("./input-test.txt"))
	fmt.Println(totalWinningsNoJoker("./input-test2.txt"))
	fmt.Println(totalWinningsNoJoker("./input.txt"))
	fmt.Println(totalWinningsWithJoker("./input-test.txt"))
	fmt.Println(totalWinningsWithJoker("./input-test2.txt"))
	fmt.Println(totalWinningsWithJoker("./input.txt"))
}
