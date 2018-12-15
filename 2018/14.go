package main

import (
	"fmt"
)

var scoreboard = []int{3, 7}

func main() {
	one_14(9)
	one_14(5)
	one_14(18)
	one_14(2018)
	one_14(909441)
	fmt.Println()
	two_14([]int{5, 1, 5, 8, 9})
	two_14([]int{0, 1, 2, 4, 5})
	two_14([]int{9, 2, 5, 1, 0})
	two_14([]int{5, 9, 4, 1, 4})
	two_14([]int{2, 6, 1, 5, 1, 6, 1, 2, 1, 3})
	two_14([]int{9, 0, 9, 4, 4, 1})
}

func two_14(k []int) {
	scoreboard = []int{3, 7}
	a, b, nReceipes := 0, 1, 0
	c := 0
	offset := 0
	for {
		if c > len(k) {
			if equals(scoreboard[len(scoreboard)-len(k):len(scoreboard)], k) {
				break
			}
			if equals(scoreboard[len(scoreboard)-len(k)-1:len(scoreboard)-1], k) {
				offset = -1
				break
			}
		}
		a, b, nReceipes = step(a, b)
		c += nReceipes
	}
	fmt.Printf("Recipes before target: %d\n", len(scoreboard)-len(k)+offset)
}

func one_14(k int) {
	scoreboard = []int{3, 7}
	a, b, nReceipes := 0, 1, 0
	c := 0
	limit := k - 2 + 10
	for c < limit {
		a, b, nReceipes = step(a, b)
		c += nReceipes
	}
	fmt.Printf("tail for %d: ", k)
	for _, r := range scoreboard[len(scoreboard)-10-(c-limit) : len(scoreboard)-(c-limit)] {
		fmt.Printf("%d", r)
	}
	fmt.Println()
}

func step(a int, b int) (int, int, int) {
	r1, r2 := makeRecipe(a, b)
	scoreboard = append(scoreboard, r1)
	newRecipes := 1
	if r2 >= 0 {
		scoreboard = append(scoreboard, r2)
		newRecipes++
	}

	return (a + scoreboard[a] + 1) % len(scoreboard),
		(b + scoreboard[b] + 1) % len(scoreboard),
		newRecipes
}

func makeRecipe(a int, b int) (int, int) {
	sum := scoreboard[a] + scoreboard[b]
	if sum < 10 {
		return sum, -1
	}
	return sum / 10, sum % 10
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
