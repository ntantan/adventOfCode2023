package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type galaxy struct {
	x int
	y int
}

type pair struct {
	path_length int
	a           galaxy
	b           galaxy
}

func main() {
	input := fileToString("input.txt")

	lines := strings.Split(input, "\n")
	rows := make([]string, 0)
	for j := range lines[0] {
		rows = append(rows, "")
		for _, line := range lines {
			rows[j] += string(line[j])
		}
	}

	fmt.Printf("lines: %q\n", lines)
	fmt.Printf("rows: %q\n", rows)

	for i := 0; i < len(lines); i++ {
		if !strings.Contains(lines[i], "#") {
			lines = slices.Insert(lines, i, lines[i])
			i++
		}
	}

	offset := 0
	for i, row := range rows {
		if !strings.Contains(row, "#") {
			pos := i + offset
			for j := range lines {
				lines[j] = lines[j][:pos] + "." + lines[j][pos:]
			}
			offset++
		}
	}

	for i, line := range lines {
		fmt.Printf("[%2d]: %q\n", i, line)
	}

	galaxies := make([]galaxy, 0)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, galaxy{x: x, y: y})
				fmt.Printf("%d\n", galaxies[len(galaxies)-1])
			}
		}
	}

	// pairs := make([]pair, 0)
	var res float64
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]
			res += math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
		}
	}

	fmt.Println(uint64(res))
}

func fileToString(path string) string {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(file)
}
