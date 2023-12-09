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
	starters := make([]node, 0)

	nodes := make(map[string]node)
	for _, link := range linking {
		token := strings.Split(link, ";")
		nodes[token[0]] = node{self: token[0], left: token[1], right: token[2]}
		if token[0][len(token[0])-1] == 'A' {
			starters = append(starters, nodes[token[0]])
		}
	}
	fmt.Printf("%q\n", starters)

	res := make([]int, 0)
	for _, starter := range starters {
		i := 0
		current := starter
		for {
			if current.self[2] == 'Z' {
				break
			}
			if instructions[i%len(instructions)] == 'R' {
				current = nodes[current.right]
			} else if instructions[i%len(instructions)] == 'L' {
				current = nodes[current.left]
			}
			i++
		}
		res = append(res, i)
		fmt.Printf("%d\n", i)
	}
	fmt.Println(LCM(res[0], res[1], res[2:]...))

	// var res, j uint64 = 0, 0
	// for {
	// 	for i := 0; i < len(starters); i++ {
	// 		if instructions[j%uint64(len(instructions))] == 'R' {
	// 			starters[i] = nodes[starters[i].right]
	// 		} else if instructions[j%uint64(len(instructions))] == 'L' {
	// 			starters[i] = nodes[starters[i].left]
	// 		}
	// 	}
	// 	res++
	// 	fmt.Printf("%q\n", starters)
	// 	if checkAllZ(starters) {
	// 		break
	// 	}
	// 	j++
	// }
	// fmt.Printf("%d\n", res)
}

// func checkAllZ(nodes []node) bool {
// 	for _, node := range nodes {
// 		if node.self[2] != 'Z' {
// 			return false
// 		}
// 	}
// 	return true
// }

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func fileToString() string {
	file, _ := os.ReadFile("input.txt")
	return string(file)
}
