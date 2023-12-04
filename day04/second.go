package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
)

func fileToString() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func checkMatch(myNums []string, winningNums []string) int {
	// fmt.Printf("%q\n", myNums)
	// fmt.Printf("%q\n", winningNums)
	count := 0
	for i := 0; i < len(myNums); i++ {
		if slices.Contains(winningNums, myNums[i]) {
			count++
		}
	}
	// fmt.Printf("%d\n", count)
	// fmt.Printf("%f\n", math.Pow(2, float64(count)))

	return count
}

func main() {
	input := fileToString()
	// fmt.Println(input)

	digitRegex, _ := regexp.Compile("\\d+")
	myNumsRegex, _ := regexp.Compile("\\:([^:]+)\\|")
	myNumsStrings := myNumsRegex.FindAllString(input, -1)
	// fmt.Printf("%q\n", myNums)
	myNums := make([][]string, 0)
	for _, myNumsString := range myNumsStrings {
		myNums = append(myNums, digitRegex.FindAllString(myNumsString, -1))
	}
	// fmt.Printf("%q\n", myNums)

	winningNumsRegex, _ := regexp.Compile("(?m)\\|([^:]+)$")
	winningNumsString := winningNumsRegex.FindAllString(input, -1)
	// fmt.Printf("%q\n", winningNumsString)
	winningNums := make([][]string, 0)
	for _, winningNumsString := range winningNumsString {
		winningNums = append(winningNums, digitRegex.FindAllString(winningNumsString, -1))
	}
	// fmt.Printf("%q\n", winningNums)

	res := 0
	array := make([]int, len(myNums))
	for i := 0; i < len(myNums); i++ {
		matchCount := checkMatch(myNums[i], winningNums[i])
		for j := i + 1; j < i+1+matchCount; j++ {
			if j < len(myNums) {
				array[j] += array[i] + 1
			}
		}
		res += array[i] + 1
		fmt.Println(res)
	}
	fmt.Println(array)
	fmt.Println(res)
}
