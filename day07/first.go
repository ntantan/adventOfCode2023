package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func fileToString() string {
	file, _ := os.ReadFile("input.txt")
	return string(file)
}

func findHandType(hand string) string {
	tab := make([]int, 6)

	for _, char := range hand {
		tab[strings.Count(hand, string(char))]++
	}

	fmt.Printf("%d\n", tab)
	if tab[5] > 0 {
		return "7" // return "Five"
	} else if tab[4] > 0 {
		return "6" // return "Four"
	} else if tab[3] > 0 && tab[2] > 0 {
		return "5" // return "FullHouse"
	} else if tab[3] > 0 {
		return "4" // return "Three"
	} else if tab[2] > 2 {
		return "3" // return "TwoPair"
	} else if tab[2] > 0 {
		return "2" // return "Pair"
	} else {
		return "1" // return "HighCards"
	}
}

func main() {
	input := fileToString()

	lines := strings.Split(input, "\n")

	for i, line := range lines {
		tokens := strings.Split(line, " ")
		lines[i] += " " + findHandType(tokens[0])
		fmt.Printf("%q\n", line)
	}

	slices.SortFunc(lines, func(a, b string) int {
		aSplit := strings.Split(a, " ")
		bSplit := strings.Split(b, " ")
		hand := cmp.Compare(aSplit[2], bSplit[2])
		if hand != 0 {
			return hand
		}

		aHand := aSplit[0]
		bHand := bSplit[0]
		for i := 0; i < len(aHand); i++ {
			if cardValues[string(aHand[i])] > cardValues[string(bHand[i])] {
				return 1
			} else if cardValues[string(aHand[i])] < cardValues[string(bHand[i])] {
				return -1
			}
		}
		return 0
	})
	fmt.Printf("%q\n", lines)

	res := 0
	for i, line := range lines {
		tokens := strings.Split(line, " ")
		atoi, _ := strconv.Atoi(tokens[1])
		res += atoi * (i + 1)
	}
	fmt.Printf("%d\n", res)
}
