package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type op = func(a, b, c int, registry map[int]int)

var ops = map[string]op{}
var opz = map[int]op{}

var addr = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] + registry[b]
}
var addi = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] + b
}
var mulr = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] * registry[b]
}
var muli = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] * b
}
var banr = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] & registry[b]
}
var bani = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] & b
}
var borr = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] | registry[b]
}
var bori = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a] | b
}
var setr = func(a, b, c int, registry map[int]int) {
	registry[c] = registry[a]
}
var seti = func(a, b, c int, registry map[int]int) {
	registry[c] = a
}
var gtir = func(a, b, c int, registry map[int]int) {
	if a > registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}
var gtri = func(a, b, c int, registry map[int]int) {
	if registry[a] > b {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}
var gtrr = func(a, b, c int, registry map[int]int) {
	if registry[a] > registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}
var eqir = func(a, b, c int, registry map[int]int) {
	if a == registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}
var eqri = func(a, b, c int, registry map[int]int) {
	if b == registry[a] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}
var eqrr = func(a, b, c int, registry map[int]int) {
	if registry[a] == registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}

func areEquals(a, b map[int]int) bool {
	equals := true
	for k := range a {
		if a[k] != b[k] {
			equals = false
			break
		}
	}
	return equals
}

var sPattern = regexp.MustCompile(`(Before|After):\s+\[(?P<a>[0-9]+),\s(?P<b>[0-9]+),\s(?P<c>[0-9]+),\s(?P<d>[0-9]+)\]`)
var opPattern = regexp.MustCompile(`(?P<a>[0-9]+)\s(?P<b>[0-9]+)\s(?P<c>[0-9]+)\s(?P<d>[0-9]+)`)

func readState(line string) map[int]int {
	match := sPattern.FindStringSubmatch(line)
	a, _ := strconv.Atoi(match[2])
	b, _ := strconv.Atoi(match[3])
	c, _ := strconv.Atoi(match[4])
	d, _ := strconv.Atoi(match[5])
	return map[int]int{0: a, 1: b, 2: c, 3: d}
}

func readOp(line string) (int, int, int, int) {
	match := opPattern.FindStringSubmatch(line)
	a, _ := strconv.Atoi(match[1])
	b, _ := strconv.Atoi(match[2])
	c, _ := strconv.Atoi(match[3])
	d, _ := strconv.Atoi(match[4])
	return a, b, c, d
}

func printRegisty(r map[int]int) {
	fmt.Printf(" %d ", r[0])
	fmt.Printf(" %d ", r[1])
	fmt.Printf(" %d ", r[2])
	fmt.Printf(" %d ", r[3])
	fmt.Println()
}

func a() {
	scanner := bufio.NewScanner(os.Stdin)
	tot := 0
	y := 0
	freq := map[int]map[string]int{}
	for scanner.Scan() {
		scanner.Scan()
		stringState := scanner.Text()
		scanner.Scan()
		o, a, b, c := readOp(scanner.Text())
		freq[o] = map[string]int{}
		scanner.Scan()
		stringFuture := scanner.Text()
		match := 0
		for k, f := range ops {
			registry := readState(stringState)
			future := readState(stringFuture)
			f(a, b, c, registry)

			if areEquals(registry, future) {
				match++
				if _, ok := freq[o][k]; !ok {
					freq[o][k] = 1
				} else {
					freq[o][k] = freq[o][k] + 1
				}
			}
		}
		if match >= 3 {
			y++
			tot++
		}
	}
	fmt.Println(tot)
}

func b() {
	scanner := bufio.NewScanner(os.Stdin)
	registry := map[int]int{0: 0, 1: 0, 2: 0, 3: 0}
	for scanner.Scan() {
		o, a, b, c := readOp(scanner.Text())
		opz[o](a, b, c, registry)
	}
	fmt.Println(registry[0])
}

func main() {
	ops["addr"] = addr
	ops["addi"] = addi
	ops["mulr"] = mulr
	ops["muli"] = muli
	ops["banr"] = banr
	ops["bani"] = bani
	ops["borr"] = borr
	ops["bori"] = bori
	ops["setr"] = setr
	ops["seti"] = seti
	ops["gtir"] = gtir
	ops["gtri"] = gtri
	ops["gtrr"] = gtrr
	ops["eqir"] = eqir
	ops["eqri"] = eqri
	ops["eqrr"] = eqrr
	//a()

	opz[5] = eqir
	opz[13] = gtrr
	opz[8] = gtri
	opz[4] = eqrr
	opz[9] = eqri
	opz[10] = gtir
	opz[15] = banr
	opz[1] = bani
	opz[3] = seti
	opz[6] = setr
	opz[12] = addr
	opz[2] = addi
	opz[11] = borr
	opz[14] = mulr
	opz[0] = muli
	opz[7] = bori
	b()
}
