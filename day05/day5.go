package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Mapper struct {
	destStart   int
	sourceStart int
	rangeLen    int
}

type Section struct {
	name    string
	mappers []Mapper
}

type Almanac struct {
	seeds    []int
	sections []Section
}

func (s Section) findDest(source int) int {
	for _, mapper := range s.mappers {
		if source >= mapper.sourceStart && source < (mapper.sourceStart+mapper.rangeLen) {
			return source + mapper.destStart - mapper.sourceStart
		}
	}
	return source
}

func (a Almanac) locationForSeed(seed int) int {
	location := seed
	for _, section := range a.sections {
		location = section.findDest(location)
	}
	return location
}

// OMG Just give me an int
func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func createAlmanac(filename string) Almanac {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	almanac := Almanac{sections: []Section{}}
	sections := strings.Split(string(b), "\n\n")
	almanac.seeds = []int{}
	for i, sec := range strings.Split(sections[0], " ") {
		if i > 0 {
			almanac.seeds = append(almanac.seeds, toInt(sec))
		}
	}

	// temperature-to-humidity map:
	// 0 69 1
	// 1 0 69
	for j := 1; j <= 7; j++ {
		newSection := Section{mappers: []Mapper{}}
		for k, sec := range strings.Split(sections[j], "\n") {
			splitLine := strings.Split(sec, " ")
			if k == 0 {
				newSection.name = splitLine[0]
			} else {
				newSection.mappers = append(newSection.mappers, Mapper{
					destStart:   toInt(splitLine[0]),
					sourceStart: toInt(splitLine[1]),
					rangeLen:    toInt(splitLine[2]),
				})
			}
		}
		almanac.sections = append(almanac.sections, newSection)
	}

	return almanac
}

func closestLocation(fileName string) int {
	a := createAlmanac(fileName)
	minLocation := 1000000000
	for _, seed := range a.seeds {
		seedLoc := a.locationForSeed(seed)
		if seedLoc < minLocation {
			minLocation = seedLoc
		}
	}
	return minLocation
}

func closestLocationByRange(fileName string) int {
	a := createAlmanac(fileName)
	minLocation := 1000000000
	for i := 0; i < len(a.seeds); i += 2 {
		for j := a.seeds[i]; j < a.seeds[i]+a.seeds[i+1]; j++ {
			seedLoc := a.locationForSeed(j)
			if seedLoc < minLocation {
				minLocation = seedLoc
			}
		}
	}
	return minLocation
}

func main() {
	fmt.Println(closestLocation("./input-test.txt"))
	fmt.Println(closestLocation("./input.txt"))
	fmt.Println(closestLocationByRange("./input-test.txt"))
	fmt.Println(closestLocationByRange("./input.txt"))
}
