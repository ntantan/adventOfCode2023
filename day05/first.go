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

func fileToString() string {
	file, err := os.ReadFile("input.txt")
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
	// seeds_to_soil := strings.Split(split[1], "\n")
	// fmt.Printf("%q\n", seeds_to_soil)
	// soil_to_fertalizer := strings.Split(split[2], "\n")
	// fmt.Printf("%q\n", soil_to_fertalizer)
	// fertilizer_to_water := strings.Split(split[3], "\n")
	// fmt.Printf("%q\n", fertilizer_to_water)
	// water_to_light := strings.Split(split[4], "\n")
	// fmt.Printf("%q\n", water_to_light)
	// temperature_to_humidity := strings.Split(split[5], "\n")
	// fmt.Printf("%q\n", temperature_to_humidity)
	// humidity_to_location := strings.Split(split[6], "\n")
	// fmt.Printf("%q\n", humidity_to_location)
	// fmt.Printf("%d\n", translate(seeds[0], seeds_to_soil))

	min := 0
	for j := 0; j < len(seeds); j++ {
		temp, _ := strconv.Atoi(seeds[j])
		for i := 1; i < 8; i++ {
			temp = translate(temp, strings.Split(split[i], "\n"))
		}
		if min == 0 || temp < min {
			min = temp
		}
		// fmt.Printf("%d\n", temp)
	}
	fmt.Printf("%d\n", min)
}
