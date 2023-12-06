package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func numWaysToBeatRecord(duration int, record int) int {
	total := 0
	for chargeTime := 0; chargeTime < duration; chargeTime++ {
		distance := chargeTime * (duration - chargeTime)
		if distance > record {
			total++
		}
	}
	return total
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func process(filename string, processor func(int, int) int) int {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(b), "\n")
	times := strings.Fields(lines[0])
	distances := strings.Fields(lines[1])
	total := 1
	for i := 1; i < len(times); i++ {
		total *= processor(toInt(times[i]), toInt(distances[i]))
	}

	return total
}

func main() {
	fmt.Println(process("./input-test.txt", numWaysToBeatRecord))
	fmt.Println(process("./input.txt", numWaysToBeatRecord))
	fmt.Println(numWaysToBeatRecord(71530, 940200))
	fmt.Println(numWaysToBeatRecord(44826981, 202107611381458))
}
