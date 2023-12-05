package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func translate(seed int, _map []string) int {
	for _, line := range _map {
		split := strings.Split(line, " ")
		destination, _ := strconv.Atoi(split[0])
		start, _ := strconv.Atoi(split[1])
		ranje, _ := strconv.Atoi(split[2])
		end := start + ranje - 1

		if seed >= start && seed <= end {
			// fmt.Printf("%q %q [seed]%d [destination]%d [start]%d [ranje]%d [end]%d\n", line, split, seed, destination, start, ranje, end)
			seed = destination + (seed - start)
			break
		}
	}

	return seed
}

func parse_seeds(seeds []string) []string {
	s := make([]string, 0)
	for i := 0; i < len(seeds); i += 2 {
		jmax, _ := strconv.Atoi(seeds[i+1])
		// fmt.Printf("%d\n", jmax)
		for j := 0; j < jmax; j++ {
			add, _ := strconv.Atoi(seeds[i])
			s = append(s, strconv.Itoa(add+j))
			// fmt.Printf("%d\n", add+j)
		}
	}
	return s
}

func fileToString() string {
	file, err := os.ReadFile("example.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

func main() {
	input := fileToString()
	split := strings.Split(input, "\n\n")

	seeds := strings.Split(split[0], " ")
	fmt.Printf("[seeds]: %q\n", seeds)

	// seeds = parse_seeds(seeds)
	// fmt.Printf("%q\n", seeds)

	min := 0
	for k := 0; k < len(seeds); k += 2 {
		jstart, _ := strconv.Atoi(seeds[k])
		jmax, _ := strconv.Atoi(seeds[k+1])
		for j := jstart; j < jstart+jmax; j++ {
			temp := j
			for i := 1; i < 8; i++ {
				temp = translate(temp, strings.Split(split[i], "\n"))
			}
			if min == 0 || temp < min {
				min = temp
			}
			fmt.Printf("Location: [%d]\n", temp)
		}
	}
	fmt.Printf("%d\n", min)
}
