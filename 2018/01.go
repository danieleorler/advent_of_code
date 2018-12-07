package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []int{}
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Println(err)
		}
		input = append(input, value)

	}
	fmt.Println(one(input))
	fmt.Println("----")
	fmt.Println(two(input))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func one(a []int) int {
	t := 0
	for _, f := range a {
		t += f
	}
	return t
}

func two(a []int) int {
	m := map[int]bool{}
	t := 0
	for {
		for _, f := range a {
			t += f
			if _, present := m[t]; present {
				return t
			} else {
				m[t] = true
			}
		}
	}
}
