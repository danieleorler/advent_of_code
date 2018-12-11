package main

import (
	"fmt"
	"log"
	"time"
)

type position struct {
	x int
	y int
}

type result struct {
	p position
	s int
	v int
}

func main() {
	start := time.Now()
	ch := make(chan result)
	for i := 2; i <= 300; i++ {
		go brute(3999, 300, i, ch)
	}

	max := result{position{0, 0}, 0, 0}
	for i := 2; i <= 300; i++ {
		t := <-ch
		if t.v > max.v {
			fmt.Printf("Current best: size %dx%d starting at [%d, %d], value: %d\n", t.s, t.s, t.p.x, t.p.y, t.v)
			max = t
		}
	}
	fmt.Printf("FINAL BEST: size %dx%d starting at [%d, %d], value: %d\n", max.s, max.s, max.p.x, max.p.y, max.v)
	log.Printf("TOTAL EXECUTION TIME: %s", time.Since(start))
}

func level(x int, y int, k int) int {
	a := ((x+10)*y + k) * (x + 10)
	return ((a % 1000) / 100) - 5
}

func brute(k int, l int, size int, channel chan result) {
	start := time.Now()
	m := 0
	p := position{0, 0}
	operations := 0
	for x := 1; x < l-(size-1); x++ {
		for y := 1; y < l-(size-1); y++ {
			tm := 0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					tm += level(x+i, y+j, k)
					operations += 1
				}
			}
			if tm > m {
				m = tm
				p = position{x, y}
			}
		}

	}
	channel <- result{p, size, m}
	log.Printf("[%dx%d] took %s executing %d operations", size, size, time.Since(start), operations)
}
