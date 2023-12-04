package main

import (
	"os"
	"strconv"
	"strings"
)

var NumberToWord = map[string]string{
	"one":       "1",
	"two":       "2",
	"three":     "3",
	"four":      "4",
	"five":      "5",
	"six":       "6",
	"seven":     "7",
	"eight":     "8",
	"nine":      "9",
	"ten":       "10",
	"eleven":    "11",
	"twelve":    "12",
	"thirteen":  "13",
	"fourteen":  "14",
	"fifteen":   "15",
	"sixteen":   "16",
	"seventeen": "17",
	"eighteen":  "18",
	"nineteen":  "19",
	"twenty":    "20",
	"thirty":    "30",
	"forty":     "40",
	"fifty":     "50",
	"sixty":     "60",
	"seventy":   "70",
	"eighty":    "80",
	"ninety":    "90",
	"hundred":   "100",
	"thousand":  "1000",
	"million":   "1000000",
	"billion":   "1000000000",
}

func findNumberToWordInString(str string) (string, bool) {
	for key := range NumberToWord {
		if strings.Contains(str, key) {
			return key, true
		}
	}
	return "", false
}

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
	arraySplit := strings.Split(string(read), "\n")
	var split []string
	for _, spl := range arraySplit {
		split = append(split, spl[:])
	}

	for _, str := range split {
		num := ""
		for start := 0; start < len(str); start++ {
			key, found := findNumberToWordInString(str[:start])
			if found {
				num += NumberToWord[key]
				break
			}
			if str[start] <= '9' && str[start] >= '0' {
				num += string(str[start])
				break
			}
		}
		for end := len(str) - 1; end >= 0; end-- {
			key, found := findNumberToWordInString(str[end:])
			if found {
				num += NumberToWord[key]
				break
			}
			if int(str[end]) <= '9' && int(str[end]) >= '0' {
				num += string(str[end])
				break
			}
		}
		intNum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		println(num, str)
		res += intNum
	}
	println(res)
}
