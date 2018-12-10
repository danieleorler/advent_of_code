package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

var pattern = regexp.MustCompile(`position=<\s*(?P<x>-?\d+),\s*(?P<y>-?\d+)>\svelocity=<\s*(?P<vx>-?\d+),\s*(?P<vy>-?\d+)>`)

type point struct {
	x int
	y int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	points := []point{}
	velocity := []point{}
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		x, err := strconv.Atoi(match[1])
		y, err := strconv.Atoi(match[2])
		vx, err := strconv.Atoi(match[3])
		vy, err := strconv.Atoi(match[4])
		points = append(points, point{x, y})
		velocity = append(velocity, point{vx, vy})

		if err != nil {
			log.Println(err)
		}
	}

	one_10_v2(points, velocity)

}

func one_10_v2(points []point, velocity []point) {
	c := 0
	minDistance := float64(^uint(0)>>1) - 1
	d := distance(points)
	for {
		for i := 0; i < len(points); i++ {
			points[i].x += velocity[i].x
			points[i].y += velocity[i].y
		}

		d = distance(points)
		if d < minDistance {
			minDistance = d
		} else {
			for i := 0; i < len(points); i++ {
				points[i].x -= velocity[i].x
				points[i].y -= velocity[i].y
			}
			break
		}
		c += 1
	}

	printMessage(points)
	fmt.Printf("\nRescue message found in %d seconds\n", c)
}

func distance(points []point) float64 {
	minx, maxx, miny, maxy := findBoundaries(points)
	return math.Sqrt(math.Pow(float64(maxx-minx), 2) + math.Pow(float64(maxy-miny), 2))
}

func findBoundaries(points []point) (int, int, int, int) {
	minx := int(^uint(0) >> 1)
	miny := int(^uint(0) >> 1)
	maxx := -minx - 1
	maxy := -miny - 1

	for _, p := range points {
		if p.x < minx {
			minx = p.x
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.y > maxy {
			maxy = p.y
		}
	}

	return minx, maxx, miny, maxy
}

func printMessage(points []point) {
	minx, maxx, miny, maxy := findBoundaries(points)
	nrows := abs(maxy) - abs(miny)
	ncol := abs(maxx) - abs(minx)
	fx := -minx
	fy := -miny

	matrix := [][]string{}
	for i := 0; i <= nrows; i++ {
		matrix = append(matrix, []string{})
		for j := 0; j <= ncol; j++ {
			matrix[i] = append(matrix[i], " ")
		}
	}

	for _, p := range points {
		if p.y+fy < len(matrix) && p.x+fx < len(matrix[0]) {
			matrix[p.y+fy][p.x+fx] = "*"
		}
	}

	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%s", cell)
		}
		fmt.Println()
	}
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}
