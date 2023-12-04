package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	read, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	res := 0
	split := strings.Split(string(read), "\n")
	for _, str := range split {
		num := ""
		for start := 0; start < len(str); start++ {
			if str[start] <= '9' && str[start] >= '0' {
				num += string(str[start])
				break
			}
		}
		for end := len(str) - 1; end >= 0; end-- {
			if int(str[end]) <= '9' && int(str[end]) >= '0' {
				num += string(str[end])
				break
			}
		}
		intNum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		res += intNum
	}
	println(res)
}
