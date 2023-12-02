package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isSeparator(c rune) bool {
	return c == ':' || c == ',' || c == ';'
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := string(data)
	// fmt.Println(input)

	lines := strings.Split(input, "\n")
	// fmt.Printf("%q\n", lines)

	var idSum int
	for _, line := range lines {
		split := strings.FieldsFunc(line, isSeparator)
		digitRegex, _ := regexp.Compile("\\d+")
		colorRegex, _ := regexp.Compile("[a-z]+")

		id, _ := strconv.Atoi(digitRegex.FindString(split[0]))
		fmt.Println(id)
		idSum += id

		for i := 1; i < len(split); i++ {
			color := colorRegex.FindString(split[i])
			num, _ := strconv.Atoi(digitRegex.FindString(split[i]))

			if (color == "red" && num > 12) || (color == "green" && num > 13) || (color == "blue" && num > 14) {
				idSum -= id
				break
			}
			println(color, num)
		}
	}

	fmt.Println(idSum)
}
