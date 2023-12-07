package main

import (
	"fmt"
	"os"
	"strconv"
)

func fileToString() string {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func main() {
	// start := time.Now()
	// file := "56717999
	// 334113513502430"
	// input := strings.Split(fileToString(), "\n")

	t := []string{"56717999"}
	d := []string{"334113513502430"}
	// fmt.Printf("%q\n%q\n", t, d)

	records := make([]int, len(t))
	for i := 0; i < len(t); i++ {
		time, _ := strconv.Atoi(t[i])
		distance, _ := strconv.Atoi(d[i])

		for start := 1; start < time+1; start++ {
			if start*(time-start) > distance {
				// fmt.Printf("[start]%d [distance]%d\n", start, distance)
				records[i]++
			}
		}
	}

	res := 1
	for _, record := range records {
		res *= record
	}
	// fmt.Printf("%d\n", records)
	fmt.Println(res)
	// elapsed := time.Since(start)
	// log.Printf("%s", elapsed)
}
