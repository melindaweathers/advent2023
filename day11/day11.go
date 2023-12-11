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

func absInt(i int) int {
	if i < 0 {
		return i * -1
	} else {
		return i
	}
}

func (galaxy Galaxy) distanceTo(otherGalaxy Galaxy) int {
	return absInt(otherGalaxy.row-galaxy.row) + absInt(otherGalaxy.col-galaxy.col)
}

func readImage(filename string) Image {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	numRows := 0
	var row Row
	image := Image{rows: []Row{}}
	for _, line := range strings.Split(string(b), "\n") {
		row = Row{cols: []rune(line)}
		numRows++
		image.rows = append(image.rows, row)
		// Append an additional blank row if this one is blank
		if row.noGalaxies() {
			row = Row{cols: []rune(line)}
			numRows++
			image.rows = append(image.rows, row)
		}
	}
	image.numRows = numRows
	image.numCols = len(row.cols)
	image = expandColumns(image)
	return image
}

func expandColumns(image Image) Image {
	newImage := Image{numRows: image.numRows, rows: []Row{}, galaxies: []Galaxy{}}
	extraCols := 0

	for r := 0; r < image.numRows; r++ {
		newImage.rows = append(newImage.rows, Row{cols: []rune{}})
	}

	for c := 0; c < image.numCols; c++ {
		hasGalaxies := false
		for r := 0; r < image.numRows; r++ {
			newRune := image.rows[r].cols[c]
			newImage.rows[r].cols = append(newImage.rows[r].cols, newRune)
			if newRune != '.' {
				newImage.galaxies = append(newImage.galaxies, Galaxy{row: r, col: c + extraCols})
				hasGalaxies = true
			}
		}
		if !hasGalaxies {
			extraCols += 1
			for r2 := 0; r2 < image.numRows; r2++ {
				newImage.rows[r2].cols = append(newImage.rows[r2].cols, '.')
			}
		}
	}
	newImage.numCols = image.numCols + extraCols
	return newImage
}

func sumDistances(filename string) int {
	image := readImage(filename)
	sum := 0

	for i := 0; i < len(image.galaxies)-1; i++ {
		for j := i + 1; j < len(image.galaxies); j++ {
			sum += image.galaxies[i].distanceTo(image.galaxies[j])
		}
	}

	return sum
}

func main() {
	fmt.Println(sumDistances("./input-test.txt"))
	fmt.Println(sumDistances("./input.txt"))
}
