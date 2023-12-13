package main

import (
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := fileToString("example.txt")
	lines := strings.Split(input, "\n")

	func(line string) {
		split := strings.Split(line, " ")
		scheme := split[0]
		strRecords := strings.Split(split[1], ",")

		records := make([]int, 0)
		//convert strings array to int arrays
		for _, rec := range strRecords {
			atoi, _ := strconv.Atoi(rec)
			records = append(records, atoi)
		}

		// count springs
		springsCount := 0
		for _, record := range records {
			springsCount += record
		}

	}(lines[0])
}

func fileToString(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}

func checkString(scheme string, records []int) bool {
	split := strings.Split(scheme, ".")

	counter := make([]int, 0)
	for _, token := range split {
		if len(token) != 0 {
			counter = append(counter, len(token))
		}
	}

	cmp := slices.Compare(counter, records)
	if cmp == 0 {
		return true
	}
	return false
}

// func main() {
// 	scheme := "###.###....###"
// 	records := []int{3, 3, 3}

// 	result := checkString(scheme, records)
// 	fmt.Println(result)
// }
