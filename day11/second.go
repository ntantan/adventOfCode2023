package main

import (
	"fmt"
	"os"
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

	// for i := 0; i < len(lines); i++ {
	// 	if !strings.Contains(lines[i], "#") {
	// 		lines = slices.Insert(lines, i, lines[i])
	// 		i++
	// 	}
	// }

	// offset := 0
	// for i, row := range rows {
	// 	if !strings.Contains(row, "#") {
	// 		pos := i + offset
	// 		for j := range lines {
	// 			lines[j] = lines[j][:pos] + "." + lines[j][pos:]
	// 		}
	// 		offset++
	// 	}
	// }

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
	var res uint64
	for i := 0; i < len(galaxies)-1; i++ {
		a := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			b := galaxies[j]

			x := 0
			if a.x < b.x {
				for k := a.x; k < b.x; k++ {
					if !strings.Contains(rows[k], "#") {
						x += 1000000
					} else {
						x++
					}
				}
			} else {
				for k := b.x; k < a.x; k++ {
					if !strings.Contains(rows[k], "#") {
						x += 1000000
					} else {
						x++
					}
				}
			}

			y := 0
			if a.y < b.y {
				for k := a.y; k < b.y; k++ {
					if !strings.Contains(lines[k], "#") {
						y += 1000000
					} else {
						y++
					}
				}
			} else {
				for k := b.y; k < a.y; k++ {
					if !strings.Contains(lines[k], "#") {
						y += 1000000
					} else {
						y++
					}
				}
			}
			// fmt.Printf("[%d][%d] %d %d\n", a, b, x, y)
			res += uint64(x + y)
			// res += math.Abs(float64(a.x-b.x)) + math.Abs(float64(a.y-b.y))
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
