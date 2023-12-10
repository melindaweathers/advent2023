package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Node struct {
	left     string
	right    string
	pathsToZ []Path
}

type Path struct {
	distance int
	start    string
	stop     string
	ptr      int
}

type Network struct {
	nodes         map[string]Node
	instructions  string
	startingNodes []string
}

func readNetwork(filename string) Network {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	network := Network{nodes: map[string]Node{}, startingNodes: []string{}}
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
			network.nodes[thisNode] = Node{pathsToZ: []Path{}, left: leftNode[1 : len(leftNode)-1], right: rightNode[:len(rightNode)-1]}
			if strings.HasSuffix(thisNode, "A") {
				network.startingNodes = append(network.startingNodes, thisNode)
			}
		}
	}
	return network
}

func countSteps(network Network, start string, suffix string, startPtr int) (int, string) {
	steps := 0
	stepName := start
	ptr := startPtr
	startNode := network.nodes[start]
	for _, path := range startNode.pathsToZ {
		if path.ptr == startPtr && path.start == start {
			return path.distance, path.stop
		}
	}
	for !strings.HasSuffix(stepName, suffix) {
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
	startNode.pathsToZ = append(startNode.pathsToZ, Path{ptr: startPtr, start: start, stop: stepName, distance: steps})
	return steps, stepName
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers []int) int {
	result := a * b / GCD(a, b)
	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i], []int{})
	}
	return result
}

func countStepsPart1(filename string) int {
	network := readNetwork(filename)
	steps, _ := countSteps(network, "AAA", "ZZZ", 0)
	return steps
}

func allDone(nodes []string) bool {
	for _, node := range nodes {
		if !strings.HasSuffix(node, "Z") {
			return false
		}
	}
	return true
}

func countGhostSteps(filename string) int {
	steps := 0
	ptr := 0
	network := readNetwork(filename)
	currentNodes := network.startingNodes
	for !allDone(currentNodes) {
		stepSize := 1
		nextNodes := []string{}
		for _, node := range currentNodes {
			stepsToZ, nodeWithZ := countSteps(network, node, "Z", ptr)
			stepSize = LCM(stepSize, stepsToZ, []int{})
			nextNodes = append(nextNodes, nodeWithZ)
		}
		steps += stepSize
		ptr = (ptr + stepSize) % len(network.instructions)
		currentNodes = nextNodes
	}
	return steps
}

func main() {
	fmt.Println(countStepsPart1("./input-test.txt"))
	fmt.Println(countStepsPart1("./input.txt"))
	fmt.Println(countGhostSteps("./input-test.txt"))
	fmt.Println(countGhostSteps("./input.txt"))
}
