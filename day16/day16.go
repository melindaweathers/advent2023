package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Beam struct {
	x         int
	y         int
	direction rune
}

type Tile struct {
	char      rune
	energized bool
}

func key(x int, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func parseKey(s string) (int, int) {
	var x, y int
	if _, err := fmt.Sscanf(s, "%d,%d", &x, &y); err != nil {
		log.Fatal(err)
	}
	return x, y
}

func readGrid(filename string) map[string]Tile {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	grid := map[string]Tile{}
	lines := strings.Split(string(b), "\n")
	energized := false
	for y, line := range lines {
		for x, char := range line {
			energized = x == 0 && y == 0
			grid[key(x, y)] = Tile{char: char, energized: energized}
		}
	}
	return grid
}

func nextKey(beam Beam) string {
	if beam.direction == 'r' {
		return key(beam.x+1, beam.y)
	} else if beam.direction == 'l' {
		return key(beam.x-1, beam.y)
	} else if beam.direction == 'u' {
		return key(beam.x, beam.y-1)
	} else if beam.direction == 'd' {
		return key(beam.x, beam.y+1)
	}
	log.Fatal("Invalid direction", beam.direction)
	return ""
}

func nextDir(char rune, startingDir rune) rune {
	var dir rune
	if char == '\\' {
		switch startingDir {
		case 'r':
			dir = 'd'
		case 'u':
			dir = 'l'
		case 'd':
			dir = 'r'
		case 'l':
			dir = 'u'
		}
	} else if char == '/' {
		switch startingDir {
		case 'r':
			dir = 'u'
		case 'u':
			dir = 'r'
		case 'd':
			dir = 'l'
		case 'l':
			dir = 'd'
		}
	} else if char == '.' {
		dir = startingDir
	} else {
		log.Fatal("Invalid dir", startingDir)
	}
	return dir
}

func countEnergized(grid map[string]Tile) int {
	sum := 0
	for _, v := range grid {
		if v.energized {
			sum++
		}
	}
	return sum
}

func printGrid(grid map[string]Tile) {
	fmt.Println("")
	for y := 0; y < 110; y++ {
		row := ""
		for x := 0; x < 110; x++ {
			tile, ok := grid[key(x, y)]
			if ok {
				if tile.energized {
					row = row + "#"
				} else {
					row = row + "."
				}
			}
		}
		if len(row) > 0 {
			fmt.Println(row)
		}
	}
}

func totalEnergized(filename string) int {
	grid := readGrid(filename)
	beams := []Beam{}
	beams = append(beams, Beam{x: -1, y: 0, direction: 'r'})
	history := map[string]bool{}

	for i := 0; len(beams) > 0; i++ {
		nextBeams := []Beam{}
		for _, beam := range beams {
			key := nextKey(beam)
			x, y := parseKey(key)
			tile, ok := grid[key]
			if !ok {
				continue
			}
			historyString := key + string(beam.direction)
			if history[historyString] {
				continue
			} else {
				history[historyString] = true
			}

			char := tile.char
			if char == '\\' || char == '/' || char == '.' {
				nextBeams = append(nextBeams, Beam{x: x, y: y, direction: nextDir(char, beam.direction)})
			} else if char == '|' {
				if beam.direction == 'u' || beam.direction == 'd' {
					nextBeams = append(nextBeams, Beam{x: x, y: y, direction: beam.direction})
				} else {
					nextBeams = append(nextBeams, Beam{x: x, y: y, direction: 'u'})
					nextBeams = append(nextBeams, Beam{x: x, y: y, direction: 'd'})
				}
			} else if char == '-' {
				if beam.direction == 'l' || beam.direction == 'r' {
					nextBeams = append(nextBeams, Beam{x: x, y: y, direction: beam.direction})
				} else {
					nextBeams = append(nextBeams, Beam{x: x, y: y, direction: 'l'})
					nextBeams = append(nextBeams, Beam{x: x, y: y, direction: 'r'})
				}
			}

			grid[key] = Tile{energized: true, char: grid[key].char}
		}
		// printGrid(grid)
		beams = nextBeams
	}

	return countEnergized(grid)
}

func main() {
	fmt.Println(totalEnergized("input-test.txt"))
	fmt.Println(totalEnergized("input-test2.txt"))
	fmt.Println(totalEnergized("input.txt"))
}
