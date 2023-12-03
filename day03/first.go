package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func checkSymAround(nums []int, SymLine [][]int, prevSymLine [][]int, nextSymLine [][]int) bool {
	for _, sym := range prevSymLine {
		if isInRange(nums, sym) {
			return true
		}
	}

	for _, sym := range SymLine {
		if isInRange(nums, sym) {
			return true
		}
	}

	for _, sym := range nextSymLine {
		if isInRange(nums, sym) {
			return true
		}
	}

	return false
}

func isInRange(a []int, b []int) bool {
	return (a[0] >= b[0] && a[0] <= b[1]) || (a[1] >= b[0] && a[1] <= b[1])
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
	// for i, digitLine := range digit {
	// 	numLine := num[i]
	// 	symLine := sym[i]
	// 	var symLinePrev [][]int
	// 	if i > 0 {
	// 		symLinePrev = sym[i-1]
	// 	}
	// 	var symLineNext [][]int
	// 	if i < len(sym) - 1 {
	// 		symLineNext = sym[i+1]
	// 	}

	// 	for j, num := range numLine {
	// 		if checkSymAround(num, symLine, symLinePrev, symLineNext) {
	// 			r,_ := strconv.Atoi(digitLine[j])
	// 			res += r
	// 		}
	// 	}

	// }

	println(res)
}
