package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Galaxy struct {
	row int
	col int
}

type Image struct {
	numRows  int
	numCols  int
	rows     []Row
	galaxies []Galaxy
}

type Row struct {
	cols []rune
}

func (r Row) noGalaxies() bool {
	for _, val := range r.cols {
		if val != '.' {
			return false
		}
	}
	return true
}

func fromTo(i1 int, i2 int) (int, int) {
	if i1 < i2 {
		return i1, i2
	} else {
		return i2, i1
	}
}

func (i Image) distanceTo(galaxy Galaxy, otherGalaxy Galaxy, expansion int) int {
	fromRow, toRow := fromTo(galaxy.row, otherGalaxy.row)
	fromCol, toCol := fromTo(galaxy.col, otherGalaxy.col)

	rowExpansions := 0
	for r := fromRow; r < toRow; r++ {
		if i.rows[r].cols[galaxy.col] == 'v' {
			rowExpansions++
		}
	}

	colExpansions := 0
	for c := fromCol; c < toCol; c++ {
		if i.rows[galaxy.row].cols[c] == 'v' {
			colExpansions++
		}
	}

	return (toRow - fromRow) + rowExpansions*(expansion-1) + (toCol - fromCol) + colExpansions*(expansion-1)
}

func readImage(filename string) Image {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	numRows := 0
	var row Row
	image := Image{rows: []Row{}, galaxies: []Galaxy{}}
	for _, line := range strings.Split(string(b), "\n") {
		row = Row{cols: []rune(line)}
		if row.noGalaxies() {
			for i := 0; i < len(row.cols); i++ {
				row.cols[i] = 'v'
			}
		}
		image.rows = append(image.rows, row)
		numRows++
	}
	image.numRows = numRows
	image.numCols = len(row.cols)
	galaxies := expandColumns(image)
	image.galaxies = galaxies
	return image
}

func expandColumns(image Image) []Galaxy {
	galaxies := []Galaxy{}
	for c := 0; c < image.numCols; c++ {
		hasGalaxies := false
		for r := 0; r < image.numRows; r++ {
			if image.rows[r].cols[c] == '#' {
				hasGalaxies = true
				galaxies = append(galaxies, Galaxy{row: r, col: c})
			}
		}
		if !hasGalaxies {
			for r := 0; r < image.numRows; r++ {
				image.rows[r].cols[c] = 'v'
			}
		}
	}
	return galaxies
}

func sumDistances(filename string, expansion int) int {
	image := readImage(filename)
	sum := 0

	for i := 0; i < len(image.galaxies)-1; i++ {
		for j := i + 1; j < len(image.galaxies); j++ {
			sum += image.distanceTo(image.galaxies[i], image.galaxies[j], expansion)
		}
	}

	return sum
}

func main() {
	fmt.Println(sumDistances("./input-test.txt", 2))
	fmt.Println(sumDistances("./input.txt", 2))
	fmt.Println(sumDistances("./input-test.txt", 10))
	fmt.Println(sumDistances("./input-test.txt", 100))
	fmt.Println(sumDistances("./input.txt", 1000000))
}
