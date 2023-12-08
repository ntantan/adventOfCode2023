package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var cardValues = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

func fileToString() string {
	file, _ := os.ReadFile("input.txt")
	return string(file)
}

func findHandType(hand string) string {
	if hand == "JJJJJ" {
		return "7"
	}

	memory := make([]byte, 0)
	tab := make([]int, 0)

	for i := 0; i < 5; i++ {
		if hand[i] != 'J' && !slices.Contains(memory, hand[i]) {
			count := strings.Count(hand, string(hand[i]))
			// tab[i] = count
			tab = append(tab, count)
			memory = append(memory, hand[i])
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(tab)))

	jCount := strings.Count(hand, "J")
	tab[0] += jCount
	fmt.Printf("sort %d\n", tab)

	if tab[0] == 5 {
		return "7" // return "Five"
	} else if tab[0] == 4 {
		return "6" // return "Four"
	} else if tab[0] == 3 && tab[1] == 2 {
		return "5" // return "FullHouse"
	} else if tab[0] == 3 {
		return "4" // return "Three"
	} else if tab[0] == 2 && tab[1] == 2 {
		return "3" // return "TwoPair"
	} else if tab[0] == 2 {
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
