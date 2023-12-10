package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Row struct {
	nums []int
}

func (r Row) allZeroes() bool {
	for _, val := range r.nums {
		if val != 0 {
			return false
		}
	}
	return true
}

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

func buildSequences(line string) []Row {
	rows := []Row{}
	row := Row{nums: []int{}}
	for _, numStr := range strings.Fields(line) {
		row.nums = append(row.nums, toInt(numStr))
	}
	rows = append(rows, row)

	for !row.allZeroes() {
		nextRow := Row{nums: []int{}}
		for i := 1; i < len(row.nums); i++ {
			nextVal := row.nums[i] - row.nums[i-1]
			nextRow.nums = append(nextRow.nums, nextVal)
		}

		rows = append(rows, nextRow)
		row = nextRow
	}
	return rows
}

func nextInHistory(line string) int {
	rows := buildSequences(line)
	newVal := 0
	rows[len(rows)-1].nums = append(rows[len(rows)-1].nums, newVal)
	for i := len(rows) - 2; i >= 0; i-- {
		newVal = rows[i].nums[len(rows[i].nums)-1] + rows[i+1].nums[len(rows[i+1].nums)-1]
		rows[i].nums = append(rows[i].nums, newVal)
	}
	return newVal
}

func nextInHistoryPart2(line string) int {
	rows := buildSequences(line)
	newVal := 0
	rows[len(rows)-1].nums = append(rows[len(rows)-1].nums, newVal)
	for i := len(rows) - 2; i >= 0; i-- {
		newVal = rows[i].nums[0] - rows[i+1].nums[0]
		rows[i].nums = append([]int{newVal}, rows[i].nums...)
	}
	return newVal
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func main() {
	fmt.Println(process("./input-test.txt", nextInHistory))
	fmt.Println(process("./input.txt", nextInHistory))
	fmt.Println(process("./input-test.txt", nextInHistoryPart2))
	fmt.Println(process("./input.txt", nextInHistoryPart2))
}
