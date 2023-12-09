package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left  string
	right string
}

type Network struct {
	nodes        map[string]Node
	instructions string
}

func readNetwork(filename string) Network {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	network := Network{nodes: map[string]Node{}}
	for i, line := range strings.Split(string(b), "\n") {
		if i == 0 {
			network.instructions = line
		} else if i >= 2 && len(line) > 1 {
			var thisNode, leftNode, rightNode string
			_, err := fmt.Sscanf(line, "%s = %s %s", &thisNode, &leftNode, &rightNode)
			if err != nil {
				panic(err)
			}
			// AAA = (BBB, CCC)
			network.nodes[thisNode] = Node{left: leftNode[1 : len(leftNode)-1], right: rightNode[:len(rightNode)-1]}
		}
	}
	return network
}

func countSteps(filename string) int {
	steps := 0
	stepName := "AAA"
	ptr := 0
	network := readNetwork(filename)
	for stepName != "ZZZ" {
		dir := network.instructions[ptr]
		if dir == 'L' {
			stepName = network.nodes[stepName].left
		} else if dir == 'R' {
			stepName = network.nodes[stepName].right
		} else {
			panic("NOT L or R")
		}
		steps++
		ptr = (ptr + 1) % len(network.instructions)
	}
	return steps
}

func main() {
	fmt.Println(countSteps("./input-test.txt"))
	fmt.Println(countSteps("./input.txt"))
}
