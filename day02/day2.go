package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	num   int
	draws []Draw
}

type Draw struct {
	red   int
	green int
	blue  int
}

// Game 9: 1 green, 5 blue; 4 blue; 2 red, 1 blue
func parseLine(line string) Game {
	outer := strings.Split(line, ": ")
	gameNum, _ := strconv.Atoi(outer[0][5:])
	lineGame := Game{
		num:   gameNum,
		draws: []Draw{},
	}
	for _, drawStr := range strings.Split(outer[1], "; ") {
		draw := Draw{}
		for _, colorStr := range strings.Split(drawStr, ", ") {
			numAndColor := strings.Split(colorStr, " ")
			numColor, _ := strconv.Atoi(numAndColor[0])
			if numAndColor[1] == "green" {
				draw.green = numColor
			} else if numAndColor[1] == "blue" {
				draw.blue = numColor
			} else if numAndColor[1] == "red" {
				draw.red = numColor
			}
		}
		lineGame.draws = append(lineGame.draws, draw)
	}
	return lineGame
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

// only 12 red cubes, 13 green cubes, and 14 blue cubes
func possibleGame(line string) int {
	game := parseLine(line)
	for _, draw := range game.draws {
		if draw.red > 12 || draw.green > 13 || draw.blue > 14 {
			return 0
		}
	}
	return game.num
}

func main() {
	fmt.Println(process("./input-test.txt", possibleGame))
	fmt.Println(process("./input.txt", possibleGame))
}
