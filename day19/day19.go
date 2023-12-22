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

type Path struct {
	instructions []Instruction
}

type Instruction struct {
	varname string
	gtlt    string
	value   int
	result  string
	negated bool
}

func toInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func readWorkflows(filename string) map[string]Workflow {
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
				instruction = Instruction{varname: parts[i], result: parts[i]}
			}
			instructions = append(instructions, instruction)
		}
		workflows[name] = Workflow{instructions: instructions}
	}
	return workflows
}

func totalValues(filename string) int {
	total := 0
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	sections := strings.Split(string(b), "\n\n")
	splitList := func(c rune) bool {
		return c == '{' || c == ',' || c == '}'
	}

	workflows := readWorkflows(filename)
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

func countCombos(filename string) int {
	workflows := readWorkflows(filename)
	pathFrontier := []Path{}

	pathFrontier = append(pathFrontier, Path{instructions: []Instruction{{result: "in"}}})

	// in{s<1351:px,qqz}
	acceptedPaths := []Path{}
	for len(pathFrontier) > 0 {
		newPaths := []Path{}
		for _, path := range pathFrontier {
			lastIns := path.instructions[len(path.instructions)-1]
			//todo is it varname or result
			if lastIns.result == "A" {
				acceptedPaths = append(acceptedPaths, path)
			} else if lastIns.result != "R" {
				workflow := workflows[lastIns.result]
				for i, ins := range workflow.instructions {
					tmp1 := make([]Instruction, len(path.instructions))
					copy(tmp1, path.instructions)
					for j := 0; j < i; j++ {
						ins1 := workflow.instructions[j]
						tmp1 = append(tmp1, Instruction{varname: ins1.varname, gtlt: ins1.gtlt, value: ins1.value, negated: true})
					}
					tmp1 = append(tmp1, Instruction{varname: ins.varname, gtlt: ins.gtlt, value: ins.value, result: ins.result})
					newPaths = append(newPaths, Path{instructions: tmp1})
				}
			}
		}
		pathFrontier = newPaths
	}

	return totalInPaths(acceptedPaths)
}

func totalInPaths(paths []Path) int {
	total := 0
	for _, path := range paths {
		total += valueInPath(path, "x") * valueInPath(path, "m") * valueInPath(path, "a") * valueInPath(path, "s")
	}

	return total
}

func valueInPath(path Path, rating string) int {
	vals := 0
	for i := 1; i <= 4000; i++ {
		if valIsValid(path, rating, i) {
			vals++
		}
	}
	return vals
}

func valIsValid(path Path, rating string, i int) bool {
	valid := true
	for _, ins := range path.instructions {
		if ins.varname == rating {
			if ins.gtlt == ">" && !ins.negated && !(i > ins.value) {
				valid = false
			} else if ins.gtlt == ">" && ins.negated && i > ins.value {
				valid = false
			} else if ins.gtlt == "<" && !ins.negated && !(i < ins.value) {
				valid = false
			} else if ins.gtlt == "<" && ins.negated && i < ins.value {
				valid = false
			}
		}
	}
	return valid
}

func main() {
	fmt.Println(totalValues("input-test.txt"))
	fmt.Println(totalValues("input.txt"))
	fmt.Println(countCombos("input-test.txt"))
	fmt.Println(countCombos("input.txt"))
}

// [{[{  0 in false} {s < 1351 px false} {a < 2006  true} {m > 2090 A false}]}
// {[{  0 in false} {s < 1351 px false} {a < 2006 qkq false} {x < 1416 A false}]}
// {[{  0 in false} {s < 1351 px false} {a < 2006  true} {m > 2090  true} {rfg  0 rfg false} {s < 537  true} {x > 2440  true} {A  0 A false}]}
// {[{  0 in false} {s < 1351  true} {qqz  0 qqz false} {s > 2770 qs false} {s > 3448 A false}]}
// {[{  0 in false} {s < 1351  true} {qqz  0 qqz false} {s > 2770  true} {m < 1801 hdj false} {m > 838 A false}]}
// {[{  0 in false} {s < 1351 px false} {a < 2006 qkq false} {x < 1416  true} {crn  0 crn false} {x > 2662 A false}]}
// {[{  0 in false} {s < 1351  true} {qqz  0 qqz false} {s > 2770 qs false} {s > 3448  true} {lnx  0 lnx false} {m > 1548 A false}]}
// {[{  0 in false} {s < 1351  true} {qqz  0 qqz false} {s > 2770 qs false} {s > 3448  true} {lnx  0 lnx false} {m > 1548  true} {A  0 A false}]}
// {[{  0 in false} {s < 1351  true} {qqz  0 qqz false} {s > 2770  true} {m < 1801 hdj false} {m > 838  true} {pv  0 pv false} {a > 1716  true} {A  0 A false}]}]
