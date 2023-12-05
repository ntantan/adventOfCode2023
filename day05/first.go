package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func translate(seed string, _map []string) int {
	s, _ := strconv.Atoi(seed)
	for _, line := range _map {
		split := strings.Split(line, " ")
		destination, _ := strconv.Atoi(split[0])
		start, _ := strconv.Atoi(split[1])
		ranje, _ := strconv.Atoi(split[2])
		end := start + ranje - 1

		if s >= start && s <= end {
			fmt.Printf("%q %q [seed]%d [destination]%d [start]%d [ranje]%d [end]%d\n", line, split, s, destination, start, ranje, end)
			s = destination + (s - start)
			break
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
	seeds_to_soil := strings.Split(split[1], "\n")
	fmt.Printf("%q\n", seeds_to_soil)
	soil_to_fertalizer := strings.Split(split[2], "\n")
	fmt.Printf("%q\n", soil_to_fertalizer)
	fertilizer_to_water := strings.Split(split[3], "\n")
	fmt.Printf("%q\n", fertilizer_to_water)
	water_to_light := strings.Split(split[4], "\n")
	fmt.Printf("%q\n", water_to_light)
	temperature_to_humidity := strings.Split(split[5], "\n")
	fmt.Printf("%q\n", temperature_to_humidity)
	humidity_to_location := strings.Split(split[6], "\n")
	fmt.Printf("%q\n", humidity_to_location)

	// fmt.Printf("%d\n", translate(seeds[0], seeds_to_soil))
	for i := 1; i < 6; i++ {
		fmt.Printf("%d\n", translate(seeds[0], strings.Split(split[i], "\n")))
	}
}
