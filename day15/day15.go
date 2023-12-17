package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func hash(s string) int {
	sum := 0
	for _, char := range s {
		sum = ((sum + int(char)) * 17) % 256
	}
	return sum
}

func sumHashes(filename string) int {
	b, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	commands := strings.Split(string(b), ",")
	sum := 0
	for _, command := range commands {
		sum += hash(command)
	}
	return sum
}

func main() {
	fmt.Println(hash("HASH"))
	fmt.Println(sumHashes("input-test.txt"))
	fmt.Println(sumHashes("input.txt"))
}
