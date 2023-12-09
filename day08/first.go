package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	self  string
	left  string
	right string
}

func main() {
	input := fileToString()

	lines := strings.Split(input, "\n")
	instructions := lines[0]
	linking := lines[1:]

	nodes := make(map[string]node)
	for _, link := range linking {
		token := strings.Split(link, ";")
		nodes[token[0]] = node{self: token[0], left: token[1], right: token[2]}
	}
	fmt.Printf("%q\n", nodes)

	i := 0
	current := nodes["AAA"]
	for {
		if current.self == "ZZZ" {
			break
		}
		if instructions[i%len(instructions)] == 'R' {
			current = nodes[current.right]
		} else if instructions[i%len(instructions)] == 'L' {
			current = nodes[current.left]
		}
		i++
	}
	fmt.Printf("%d\n", i)
}

func fileToString() string {
	file, _ := os.ReadFile("input.txt")
	return string(file)
}
