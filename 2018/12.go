package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var statePattern = regexp.MustCompile(`initial state:\s(?P<x>.*)`)
var rulePattern = regexp.MustCompile(`(?P<scenario>[.#]+)\s=>(?P<result>\s[.#]{1})`)

type rule struct {
	scenario string
	result   string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	initialState := statePattern.FindStringSubmatch(scanner.Text())[1]
	scanner.Scan()
	rules := map[string]bool{}
	for scanner.Scan() {
		match := rulePattern.FindStringSubmatch(scanner.Text())
		//rule := rule{match[1], match[2]}
		if strings.TrimSpace(match[2]) == "#" {
			rules[match[1]] = true
		} else {
			rules[match[1]] = false
		}

	}
	one_12(initialState, rules)
	two_12(initialState, rules)
}

func one_12(initialState string, rules map[string]bool) {
	leftPad := "..."
	rightPad := "..."
	status := leftPad + initialState + rightPad
	c := 0
	for g := 1; g <= 20; g++ {
		nextStatus := ""
		c = 0
		for i := 2; i < len(status)-2; i++ {
			if willSpread, isThere := rules[status[i-2:i+3]]; willSpread && isThere {
				nextStatus += "#"
				c += i - len(leftPad)
			} else {
				nextStatus += "."
			}
		}
		status = ".." + nextStatus + rightPad
	}
	log.Printf("Result one: %d", c)
}

func two_12(initialState string, rules map[string]bool) {
	leftPad := "..."
	rightPad := "..."
	status := leftPad + initialState + rightPad
	c := 0
	g := 1
	prev := 0
	for {
		nextStatus := ""
		c = 0
		for i := 2; i < len(status)-2; i++ {
			if willSpread, isThere := rules[status[i-2:i+3]]; willSpread && isThere {
				nextStatus += "#"
				c += i - len(leftPad)
			} else {
				nextStatus += "."
			}
		}

		a := strings.TrimLeft(".."+nextStatus+rightPad, ".")
		b := strings.TrimLeft(status, ".")
		if len(a) == len(b) && a == b {
			break
		}
		status = ".." + nextStatus + rightPad
		g += 1
		prev = c
	}
	log.Printf("Current: %d, Diff: %d, Generation: %d, Result two: %d", c, c-prev, g, c+(50000000000-g)*(c-prev))
}
