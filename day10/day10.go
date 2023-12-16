package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Tile struct {
	char rune
	dist int
	x    int
	y    int
}

type Field struct {
	tiles    map[string]*Tile
	frontier []*Tile
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

func (f Field) northOf(t Tile) *Tile {
	tilePtr, ok := f.tiles[key(t.x, t.y-1)]
	if ok && ((*tilePtr).char == '7' || (*tilePtr).char == 'F' || (*tilePtr).char == '|') {
		return tilePtr
	} else {
		return &Tile{char: '.'}
	}
}

func (f Field) eastOf(t Tile) *Tile {
	tilePtr, ok := f.tiles[key(t.x+1, t.y)]
	if ok && ((*tilePtr).char == 'J' || (*tilePtr).char == '-' || (*tilePtr).char == '7') {
		return tilePtr
	} else {
		return &Tile{char: '.'}
	}
}
func (f Field) westOf(t Tile) *Tile {
	tilePtr, ok := f.tiles[key(t.x-1, t.y)]
	if ok && ((*tilePtr).char == 'L' || (*tilePtr).char == 'F' || (*tilePtr).char == '-') {
		return tilePtr
	} else {
		return &Tile{char: '.'}
	}
}
func (f Field) southOf(t Tile) *Tile {
	tilePtr, ok := f.tiles[key(t.x, t.y+1)]
	if ok && ((*tilePtr).char == 'L' || (*tilePtr).char == 'J' || (*tilePtr).char == '|') {
		return tilePtr
	} else {
		return &Tile{char: '.'}
	}
}

func appendIfTile(tiles []*Tile, t *Tile) []*Tile {
	if t.char == '.' {
		return tiles
	} else {
		return append(tiles, t)
	}
}

func buildField(filename string) Field {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	field := Field{tiles: map[string]*Tile{}, frontier: []*Tile{}}
	for y, line := range strings.Split(string(b), "\n") {
		for x, char := range line {
			tile := Tile{char: char, dist: -1, x: x, y: y}
			field.tiles[key(x, y)] = &tile
			if char == 'S' {
				field.frontier = append(field.frontier, &tile)
				tile.dist = 0
			}
		}
	}
	return field
}

func (f Field) print() {
	// fmt.Println("TILES:")
	// for _, t := range f.tiles {
	// 	(*t).print()
	// }
	fmt.Println("FRONTIER:")
	for _, t := range f.frontier {
		(*t).print()
	}
	fmt.Println("------")
}

func (f Tile) print() {
	fmt.Println(string(f.char), ":", f.dist, f.x, f.y)
}

func (f Field) findNeighbors(t Tile) []*Tile {
	initialNeighbors := []*Tile{}
	if t.char == '|' {
		initialNeighbors = appendIfTile(initialNeighbors, f.southOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.northOf(t))
	} else if t.char == '-' {
		initialNeighbors = appendIfTile(initialNeighbors, f.westOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.eastOf(t))
	} else if t.char == 'L' {
		initialNeighbors = appendIfTile(initialNeighbors, f.northOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.eastOf(t))
	} else if t.char == 'J' {
		initialNeighbors = appendIfTile(initialNeighbors, f.northOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.westOf(t))
	} else if t.char == '7' {
		initialNeighbors = appendIfTile(initialNeighbors, f.southOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.westOf(t))
	} else if t.char == 'F' {
		initialNeighbors = appendIfTile(initialNeighbors, f.southOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.eastOf(t))
	} else if t.char == 'S' {
		initialNeighbors = appendIfTile(initialNeighbors, f.southOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.eastOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.westOf(t))
		initialNeighbors = appendIfTile(initialNeighbors, f.northOf(t))
	}
	return initialNeighbors
}

func findLongest(filename string) int {
	field := buildField(filename)
	dist := 0

	for len(field.frontier) > 0 {
		// field.print()
		nextFrontier := []*Tile{}
		for _, tile := range field.frontier {
			neighbors := field.findNeighbors(*tile)
			for _, n := range neighbors {
				if n.dist == -1 {
					n.dist = dist + 1
					nextFrontier = append(nextFrontier, n)
				}
			}

		}
		dist++
		field.frontier = nextFrontier
	}
	return dist - 1
}

func main() {
	fmt.Println(findLongest("./input-test.txt"))
	fmt.Println(findLongest("./input-test2.txt"))
	fmt.Println(findLongest("./input.txt"))
}
