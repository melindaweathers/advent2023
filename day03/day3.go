package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Schematic struct {
	max_i  int
	max_j  int
	engine map[string]rune
	gears  map[string][]int
}

func (s Schematic) isNum(i int, j int) bool {
	num := s.engine[key(i, j)] - '0'
	return num >= 0 && num <= 9
}

func (s Schematic) isSymbol(i int, j int) bool {
	return !s.isNum(i, j) && !(s.val(i, j) == '.')
}

func (s Schematic) val(i int, j int) rune {
	runeVal, ok := s.engine[key(i, j)]
	if !ok {
		runeVal = '.'
	}
	return runeVal
}

func process(filename string, processor func(Schematic) int) int {
	schematic := Schematic{}
	schematic.engine = map[string]rune{}
	schematic.gears = map[string][]int{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	max_j := 0
	for scanner.Scan() {
		for j, character := range scanner.Text() {
			schematic.engine[key(i, j)] = character
			max_j = j
		}
		i += 1
	}
	schematic.max_i = i - 1
	schematic.max_j = max_j

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return processor(schematic)
}

func key(i int, j int) string {
	return fmt.Sprintf("%d,%d", i, j)
}

func isSymbolAdjacent(s Schematic, i int, jStart int, jEnd int, val int) bool {
	foundSymbol := false
	for iTest := i - 1; iTest <= i+1; iTest++ {
		for jTest := jStart - 1; jTest <= jEnd+1; jTest++ {
			if s.isSymbol(iTest, jTest) {
				foundSymbol = true
			}
			if s.val(iTest, jTest) == '*' {
				loc := key(iTest, jTest)
				s.gears[loc] = append(s.gears[loc], val)
			}
		}
	}
	return foundSymbol
}

func processSchematic(s Schematic) int {
	sum := 0
	insideNumber := false
	jStart := -1
	jEnd := -1
	partNumber := ""
	for i := 0; i <= int(s.max_i); i++ {
		for j := 0; j <= int(s.max_j); j++ {
			if !insideNumber && s.isNum(i, j) {
				jStart = j
				insideNumber = true
			}
			if insideNumber {
				partNumber = partNumber + string(s.val(i, j))
				if !s.isNum(i, j+1) {
					jEnd = j
					insideNumber = false
					partNumberNum, _ := strconv.Atoi(partNumber)
					if isSymbolAdjacent(s, i, jStart, jEnd, partNumberNum) {
						sum += partNumberNum
					}
					partNumber = ""
				}
			}
		}
	}
	return sum
}

func sumPartNumbers(s Schematic) int {
	return processSchematic(s)
}

func sumGearRatios(s Schematic) int {
	_ = processSchematic(s)
	sum := 0
	for _, v := range s.gears {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}

func main() {
	fmt.Println(process("./input-test.txt", sumPartNumbers))
	fmt.Println(process("./input.txt", sumPartNumbers))
	fmt.Println(process("./input-test.txt", sumGearRatios))
	fmt.Println(process("./input.txt", sumGearRatios))
}
