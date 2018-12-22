package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type acre struct {
	x int
	y int
}

func readArea(scanner *bufio.Scanner) (map[acre]rune, acre, acre) {
	y := 0
	maxx := 0
	area := map[acre]rune{}
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for x, r := range line {
			area[acre{x, y}] = r
			maxx = x
		}
		y++
	}

	return area, acre{0, 0}, acre{maxx, y}
}

func countSorroundings(area map[acre]rune, current acre) map[rune]int {
	f := map[rune]int{}
	for y := current.y - 1; y <= current.y+1; y++ {
		for x := current.x - 1; x <= current.x+1; x++ {
			if v, b := area[acre{x, y}]; b && (x != current.x || y != current.y) {
				if _, there := f[v]; !there {
					f[v] = 0
				}
				f[v] = f[v] + 1
			}
		}
	}
	return f
}

func transform(v rune, f map[rune]int) rune {
	if v == '.' && f['|'] >= 3 {
		return '|'
	}
	if v == '|' && f['#'] >= 3 {
		return '#'
	}
	if v == '#' && (f['#'] < 1 || f['|'] < 1) {
		return '.'
	}
	return v
}

func tick(area map[acre]rune, min acre, max acre) map[acre]rune {
	newArea := map[acre]rune{}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if v, b := area[acre{x, y}]; b {
				f := countSorroundings(area, acre{x, y})
				newArea[acre{x, y}] = transform(v, f)
			}
		}
	}
	return newArea
}

func value(area map[acre]rune, min acre, max acre) int {
	f := map[rune]int{}
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if v, b := area[acre{x, y}]; b {
				if _, there := f[v]; !there {
					f[v] = 0
				}
				f[v] = f[v] + 1
			}
		}
	}
	return f['|'] * f['#']
}

func main() {
	area, min, max := readArea(bufio.NewScanner(os.Stdin))
	set := map[int]int{}
	for i := 0; i < 10; i++ {
		area = tick(area, min, max)
	}
	fmt.Printf("Resource value: %d\n", value(area, min, max))
	// part two: values repeats (in cycle of 28) after tick 539
	// 539 + x * 28 = 10^9  --> x = 35714266 adn change
	// 10^9 - 539 + 35714266 * 28 --> gives the index-1 of value in the 28-cycle
}
