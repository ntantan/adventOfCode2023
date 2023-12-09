package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := fileToString("input.txt")
	lines := strings.Split(input, "\n")
	splittedLines := make([][]string, len(lines))
	for _, line := range lines {
		splittedLines = append(splittedLines, strings.Split(line, " "))
	}
	fmt.Printf("%q\n", splittedLines)

	res := 0
	for _, line := range splittedLines {
		tabs := make([][]int, 1)
		for _, str := range line {
			atoi, _ := strconv.Atoi(str)
			tabs[0] = append(tabs[0], atoi)
		}

		for i := 1; i < len(tabs[0]); i++ {
			if !isNotFullZeroTab(tabs[i-1]) {
				value := 0
				for k := len(tabs) - 1; k > 0; k-- {
					// fmt.Printf("%d %d\n", value, tabs[k-1][len(tabs[k])-1])
					value = value + tabs[k-1][len(tabs[k-1])-1]
				}
				fmt.Printf("%d\n", value)
				res += value
				break
			}
			newTab := make([]int, 0)
			for j := 0; j < len(tabs[i-1])-1; j++ {
				newTab = append(newTab, tabs[i-1][j+1]-tabs[i-1][j])
			}
			tabs = append(tabs, newTab)
			// fmt.Printf("%d\n", tabs)
		}
	}

	fmt.Printf("%d\n", res)

}

func isNotFullZeroTab(tab []int) bool {
	return slices.ContainsFunc(tab, func(n int) bool {
		return n != 0
	})
}

func fileToString(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}
