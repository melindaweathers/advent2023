package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Workflow struct {
	instructions []Instruction
}

type Instruction struct {
	varname string
	gtlt    string
	value   int
	result  string
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

// qqz{s>2770:qs,m<1801:hdj,R}
// gd{a>3333:R,R}

// {x=787,m=2655,a=1222,s=2876}
// {x=1679,m=44,a=2067,s=496}
func totalValues(filename string) int {
	total := 0
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	sections := strings.Split(string(b), "\n\n")
	workflows := map[string]Workflow{}
	splitList := func(c rune) bool {
		return c == '{' || c == ',' || c == '}'
	}
	for _, workflowLine := range strings.Split(sections[0], "\n") {
		parts := strings.FieldsFunc(workflowLine, splitList)
		name := parts[0]
		instructions := []Instruction{}
		var instruction Instruction
		for i := 1; i < len(parts); i++ {
			gtlt := strings.IndexAny(parts[i], "<>")
			if gtlt > 0 {
				col := strings.IndexAny(parts[i], ":")
				instruction = Instruction{varname: parts[i][0:gtlt], gtlt: parts[i][gtlt : gtlt+1], value: toInt(parts[i][gtlt+1 : col]), result: parts[i][col+1:]}
			} else {
				instruction = Instruction{varname: parts[i]}
			}
			instructions = append(instructions, instruction)
		}
		workflows[name] = Workflow{instructions: instructions}
	}

	for _, ratingsLine := range strings.Split(sections[1], "\n") {
		parts := strings.FieldsFunc(ratingsLine, splitList)
		ratings := map[string]int{}
		for _, part := range parts {
			ratingsParts := strings.Split(part, "=")
			ratings[ratingsParts[0]] = toInt(ratingsParts[1])
		}

		total += calc("in", ratings, workflows)
	}
	return total
}

func ratingsValue(r map[string]int) int {
	total := 0
	for _, v := range r {
		total += v
	}
	return total
}

func calc(start string, ratings map[string]int, workflows map[string]Workflow) int {
	if start == "A" {
		return ratingsValue(ratings)
	} else if start == "R" {
		return 0
	}
	for _, ins := range workflows[start].instructions {
		if len(ins.gtlt) > 0 {
			if ins.gtlt == "<" {
				if ratings[ins.varname] < ins.value {
					return calc(ins.result, ratings, workflows)
				}
			} else {
				if ratings[ins.varname] > ins.value {
					return calc(ins.result, ratings, workflows)
				}
			}
		} else if ins.varname == "A" {
			return ratingsValue(ratings)
		} else if ins.varname == "R" {
			return 0
		} else {
			return calc(ins.varname, ratings, workflows)
		}
	}

	log.Fatal("What happened")
	return 0
}

func main() {
	fmt.Println(totalValues("input-test.txt"))
	fmt.Println(totalValues("input.txt"))
}
