package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkAroundSym(sym []int, numLine [][]int, numLinePrev [][]int, numLineNext [][]int, digitLine []string, digitLinePrev []string, digitLineNext []string) int {
	var adj []int
	for i, num := range numLine {
		if isInRange(sym, num) {
			atoi, _ := strconv.Atoi(digitLine[i])
			adj = append(adj, atoi)
		}
	}

	for i, num := range numLinePrev {
		if isInRange(sym, num) {
			atoi, _ := strconv.Atoi(digitLinePrev[i])
			adj = append(adj, atoi)
		}
	}

	for i, num := range numLineNext {
		if isInRange(sym, num) {
			atoi, _ := strconv.Atoi(digitLineNext[i])
			adj = append(adj, atoi)
		}
	}

	if len(adj) >= 2 {
		fmt.Printf("[%d][%d]-", adj[0], adj[1])
		return adj[0] * adj[1]
	}
	return 0
}

func isInRange(a []int, b []int) bool {
	return (a[0] >= b[0] && a[0] <= b[1]) || (a[1] >= b[0] && a[1] <= b[1]) || (b[0] >= a[0] && b[0] <= a[1]) || (b[1] >= a[0] && b[1] <= a[1])
}

func fileToString() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func main() {
	input := fileToString()

	split := strings.Split(input, "\n")
	fmt.Printf("%q\n", split)

	digitRegex, _ := regexp.Compile("\\d+")
	symbolRegex, _ := regexp.Compile("\\*+")
	num, sym, digit := make([][][]int, 0), make([][][]int, 0), make([][]string, 0)
	for i, line := range split {
		digitLine := digitRegex.FindAllString(line, -1)
		fmt.Printf("digit[%d] : %q\n", i, digitLine)
		digit = append(digit, digitLine)

		numLine := digitRegex.FindAllIndex([]byte(line), -1)
		for _, num := range numLine {
			num[1]--
		}
		fmt.Printf("num[%d] : %d\n", i, numLine)
		num = append(num, numLine)

		symLine := symbolRegex.FindAllIndex([]byte(line), -1)
		for _, sym := range symLine {
			if sym[0] > 0 {
				sym[0]--
			}
		}
		fmt.Printf("sym[%d] : %d\n", i, symLine)
		sym = append(sym, symLine)
	}

	res := 0
	for i, digitLine := range digit {
		symLine := sym[i]
		numLine := num[i]
		var numLinePrev [][]int
		var digitLinePrev []string
		if i > 0 {
			numLinePrev = num[i-1]
			digitLinePrev = digit[i-1]
		}
		var numLineNext [][]int
		var digitLineNext []string
		if i < len(num)-1 {
			numLineNext = num[i+1]
			digitLineNext = digit[i+1]
		}

		for _, sym := range symLine {
			res += checkAroundSym(sym, numLine, numLinePrev, numLineNext, digitLine, digitLinePrev, digitLineNext)
			fmt.Printf("|%d|", res)
		}
		fmt.Println(i)

	}

	println(res)
}
