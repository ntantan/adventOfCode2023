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

	var cubesSum int
	for _, line := range lines {
		split := strings.FieldsFunc(line, isSeparator)
		digitRegex, _ := regexp.Compile("\\d+")
		colorRegex, _ := regexp.Compile("[a-z]+")

		id, _ := strconv.Atoi(digitRegex.FindString(split[0]))
		fmt.Println(id)

		red, blue, green := 0, 0, 0
		for i := 1; i < len(split); i++ {
			color := colorRegex.FindString(split[i])
			num, _ := strconv.Atoi(digitRegex.FindString(split[i]))
			// println(color, num)

			if color == "red" && num > red {
				red = num
			} else if color == "blue" && num > blue {
				blue = num
			} else if color == "green" && num > green {
				green = num
			}
			println(red, blue, green)
		}
		cubesSum += red * blue * green
		println(red * blue * green)
	}

	fmt.Println(cubesSum)
}
