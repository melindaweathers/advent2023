package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Lens struct {
	label       string
	focalLength int
}

type Box struct {
	lenses []Lens
}

func (b Box) lensIndex(label string) int {
	for i, lens := range b.lenses {
		if lens.label == label {
			return i
		}
	}
	return -1
}

func (b Box) remove(label string) Box {
	index := b.lensIndex(label)
	if index >= 0 {
		b.lenses = slices.Delete(b.lenses, index, index+1)
	}
	return b
}

func (b Box) add(label string, power int) Box {
	index := b.lensIndex(label)
	newLens := Lens{label: label, focalLength: power}
	if index >= 0 {
		b.lenses[index] = newLens
	} else {
		b.lenses = append(b.lenses, newLens)
	}
	return b
}

func hash(s string) int {
	sum := 0
	for _, char := range s {
		sum = ((sum + int(char)) * 17) % 256
	}
	return sum
}

func readCommands(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(b), ",")
}

func sumHashes(filename string) int {
	commands := readCommands(filename)
	sum := 0
	for _, command := range commands {
		sum += hash(command)
	}
	return sum
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func findFocalLength(filename string) int {
	commands := readCommands(filename)
	boxes := make([]Box, 256)
	for i := 0; i < 256; i++ {
		boxes[i] = Box{lenses: []Lens{}}
	}
	var pieces []string
	// rn=1,cm-
	for _, command := range commands {
		var label string
		if strings.HasSuffix(command, "-") {
			label = command[:len(command)-1]
			boxes[hash(label)] = boxes[hash(label)].remove(label)
		} else {
			pieces = strings.Split(command, "=")
			label = pieces[0]
			boxes[hash(label)] = boxes[hash(label)].add(label, toInt(pieces[1]))
		}
	}

	focusingPower := 0
	for i := 0; i < 256; i++ {
		box := boxes[i]
		for l, lens := range box.lenses {
			focusingPower += (1 + i) * (1 + l) * lens.focalLength
		}
	}
	return focusingPower
}

func main() {
	fmt.Println(hash("HASH"))
	fmt.Println(sumHashes("input-test.txt"))
	fmt.Println(sumHashes("input.txt"))
	fmt.Println(findFocalLength("input-test.txt"))
	fmt.Println(findFocalLength("input.txt"))
}
