package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var clayPattern = regexp.MustCompile(`(?P<a>x|y)=(?P<x>[0-9]+),\s(?P<b>x|y)=(?P<y1>[0-9]+)\.\.(?P<y2>[0-9]+)`)

type coordinate struct {
	x int
	y int
}

var emptyCoord = coordinate{-1, -1}
var startCountingFromY int

func readScan(scanner *bufio.Scanner) (map[coordinate]rune, coordinate, coordinate) {
	matrix := map[coordinate]rune{}
	maxx := math.MinInt64
	minx := math.MaxInt64
	maxy := math.MinInt64
	miny := math.MaxInt64
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		match := clayPattern.FindStringSubmatch(line)

		lower, _ := strconv.Atoi(match[4])
		upper, _ := strconv.Atoi(match[5])
		if match[1] == "x" {
			x, _ := strconv.Atoi(match[2])
			if x > maxx {
				maxx = x
			}
			if x < minx {
				minx = x
			}

			for y := lower; y <= upper; y++ {
				if y > maxy {
					maxy = y
				}
				if y < miny {
					miny = y
				}
				matrix[coordinate{x, y}] = '#'
			}
		}
		if match[1] == "y" {
			y, _ := strconv.Atoi(match[2])
			if y > maxy {
				maxy = y
			}
			if y < miny {
				miny = y
			}
			for x := lower; x <= upper; x++ {
				if x > maxx {
					maxx = x
				}
				if x < minx {
					minx = x
				}
				matrix[coordinate{x, y}] = '#'
			}
		}
	}
	startCountingFromY = miny
	return matrix, coordinate{minx - 1, 0}, coordinate{maxx + 1, maxy + 1}
}

func drawScan(scan map[coordinate]rune, min coordinate, max coordinate) {
	fmt.Printf("\033[0;0H")
	for y := min.y; y <= max.y; y++ {
		for x := min.x; x <= max.x; x++ {
			if v, b := scan[coordinate{x, y}]; b {
				fmt.Printf(" %c ", v)
			} else if x == 500 && y == 0 {
				fmt.Printf(" %c ", '+')
			} else {
				fmt.Printf(" %c ", '.')
			}
		}
		fmt.Println()
	}
}

func findOverflow(scan map[coordinate]rune, current coordinate, min coordinate, max coordinate) (action string, left coordinate, right coordinate) {
	y := current.y
	action = "NOTHING"
	for x := current.x; x >= min.x; x-- {
		material, isOccupied := scan[coordinate{x, current.y}]
		materialBelow, isBelowOccupied := scan[coordinate{x, current.y + 1}]

		if material == '#' {
			action = "FILL"
			left = coordinate{x, y}
			break
		}
		if (!isOccupied || material == '|') && (!isBelowOccupied || materialBelow == '|') {
			action = "OVERFLOW_LEFT"
			left = coordinate{x, y}
			break
		}
	}

	for x := current.x; x <= max.x; x++ {
		material, isOccupied := scan[coordinate{x, current.y}]
		materialBelow, isBelowOccupied := scan[coordinate{x, current.y + 1}]

		if material == '#' {
			if action != "OVERFLOW_LEFT" {
				action = "FILL"
			}
			right = coordinate{x, y}
			break
		}
		if (!isOccupied || material == '|') && (!isBelowOccupied || materialBelow == '|') {
			if action == "OVERFLOW_LEFT" {
				action = "OVERFLOW_BOTH"
			} else {
				action = "OVERFLOW_RIGHT"
			}
			right = coordinate{x, y}
			break
		}
	}

	return action, left, right
}

func fill(scan map[coordinate]rune, from coordinate, to coordinate, material rune) map[coordinate]rune {
	for x := from.x + 1; x < to.x; x++ {
		scan[coordinate{x, from.y}] = material
	}
	return scan
}

func tick(scan map[coordinate]rune, spring coordinate, min coordinate, max coordinate) map[coordinate]rune {
	down := coordinate{spring.x, spring.y + 1}
	materialBelow, downBlocked := scan[down]
	if down.y >= stop {
		return scan
	}

	for !downBlocked && spring.y < stop {
		scan[down] = '|'
		spring = down
		down = coordinate{spring.x, spring.y + 1}
		materialBelow, downBlocked = scan[down]
	}
	if materialBelow == '|' {
		return scan
	}
	if spring.y >= stop {
		return scan
	}

	action, left, right := findOverflow(scan, spring, min, max)
	for action == "FILL" {
		scan = fill(scan, left, right, '~')
		spring = coordinate{spring.x, spring.y - 1}
		action, left, right = findOverflow(scan, spring, min, max)
	}

	if action == "OVERFLOW_RIGHT" {
		scan = fill(scan, left, coordinate{right.x + 1, right.y}, '|')
		return tick(scan, right, min, max)
	}

	if action == "OVERFLOW_LEFT" {
		scan = fill(scan, coordinate{left.x - 1, left.y}, right, '|')
		return tick(scan, left, min, max)
	}

	if action == "OVERFLOW_BOTH" {
		scan = fill(scan, coordinate{left.x - 1, left.y}, coordinate{right.x + 1, right.y}, '|')
		return tick(tick(scan, left, min, max), right, min, max)
	}

	if action == "NOTHING" {
		return scan
	}

	return scan
}

var stop int

func main() {
	scan, min, max := readScan(bufio.NewScanner(os.Stdin))
	stop = max.y - 1
	spring := coordinate{500, 0}
	tick(scan, spring, min, coordinate{max.x, max.y + 5})
	water := 0
	for k, v := range scan {
		if k.y >= startCountingFromY && (v == '|' || v == '~') {
			water++
		}
	}
	fmt.Println(water)
}
