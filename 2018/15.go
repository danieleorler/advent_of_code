package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

type coordinate struct {
	x int
	y int
}

type tail interface {
	at() coordinate
	render() rune
}

type player interface {
	at() coordinate
	move(to coordinate)
	render() rune
	hpLeft() int
	enemy() rune
	suffer()
}

type elf struct {
	coord coordinate
	hp    int
}

func (e elf) render() rune {
	return 'E'
}
func (e elf) at() coordinate {
	return e.coord
}
func (e elf) move(to coordinate) {
	e.coord = to
}
func (e elf) hpLeft() int {
	return e.hp
}
func (e elf) enemy() rune {
	return 'G'
}
func (e elf) suffer() {
	e.hp -= 3
}

type goblin struct {
	coord coordinate
	hp    int
}

func (g goblin) render() rune {
	return 'G'
}
func (g goblin) at() coordinate {
	return g.coord
}
func (g goblin) move(to coordinate) {
	g.coord = to
}
func (g goblin) hpLeft() int {
	return g.hp
}
func (g goblin) enemy() rune {
	return 'E'
}
func (g goblin) suffer() {
	g.hp -= 3
}

type wall struct {
	coord coordinate
}

func (w wall) render() rune {
	return '#'
}
func (w wall) at() coordinate {
	return w.coord
}

type empty struct {
	coord coordinate
}

func (e empty) render() rune {
	return '.'
}
func (e empty) at() coordinate {
	return e.coord
}

type Players []*player

func (list Players) Len() int      { return len(list) }
func (list Players) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list Players) Less(i, j int) bool {
	a := *list[i]
	b := *list[j]
	return a.at().y < b.at().y || a.at().y == b.at().y && a.at().x < b.at().x
}

type Tails []player

func (list Tails) Len() int      { return len(list) }
func (list Tails) Swap(i, j int) { list[i], list[j] = list[j], list[i] }
func (list Tails) Less(i, j int) bool {
	a := list[i]
	b := list[j]
	return a.at().y < b.at().y || a.at().y == b.at().y && a.at().x < b.at().x
}

func createTail(a rune, x, y int) tail {
	c := coordinate{x, y}
	switch a {
	case 'E':
		return elf{c, 200}
	case 'G':
		return goblin{c, 200}
	case '#':
		return wall{c}
	default:
		return empty{c}
	}
}

func dist(a, b coordinate) int {
	d := math.Abs(float64(a.x)-float64(b.x)) + math.Abs(float64(a.y)-float64(b.y))
	return int(d)
}

func readInput(scanner *bufio.Scanner) []string {
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}
	return lines
}

func readMap(lines []string) ([][]tail, Players) {
	game := [][]tail{}
	players := Players{}
	y := 0
	for _, line := range lines {
		game = append(game, []tail{})
		for x, r := range line {
			game[y] = append(game[y], createTail(r, x, y))
			if p, b := game[y][x].(player); b {
				players = append(players, &p)
			}
		}
		y++
	}
	return game, players
}

func findEnemies(p player, game [][]tail) []player {
	enemies := Tails{}
	for y := 0; y < len(game); y++ {
		for x := 0; x < len(game[y]); x++ {
			if p.enemy() == game[y][x].render() && game[y][x].(player).hpLeft() > 0 {
				enemies = append(enemies, game[y][x].(player))
			}
		}
	}
	sort.Sort(enemies)
	return enemies
}
func findTargets(game [][]tail, e []player) []empty {
	targets := []empty{}
	for _, enemy := range e {
		x := enemy.at().x
		y := enemy.at().y
		if y-1 < len(game[0]) && y < len(game) {
			if v, b := game[y-1][x].(empty); b {
				targets = append(targets, v)
			}
		}
		if x-1 < len(game[0]) && y < len(game) {
			if v, b := game[y][x-1].(empty); b {
				targets = append(targets, v)
			}
		}
		if x+1 < len(game[0]) && y < len(game) {
			if v, b := game[y][x+1].(empty); b {
				targets = append(targets, v)
			}
		}
		if y+1 < len(game[0]) && y < len(game) {
			if v, b := game[y+1][x].(empty); b {
				targets = append(targets, v)
			}
		}
	}
	return targets
}

func findClosest(p player, game [][]tail, r []empty) (bool, empty) {
	var target empty
	closest := math.MaxInt64
	for _, c := range r {
		if reachable, path := astar(p, c, game); reachable {
			d := len(path)
			if d < closest {
				closest = d
				target = c
			}
			if d == closest && (c.at().y < target.at().y || (c.at().y == target.at().y && c.at().x < target.at().x)) {
				closest = d
				target = c
			}
		}
	}
	return closest != math.MaxInt64, target
}

func nextStep(p player, game [][]tail, t empty) coordinate {
	minLen := math.MaxInt64
	next := coordinate{-1, -1}
	targets := []empty{}
	x := p.at().x
	y := p.at().y
	if y-1 < len(game[0]) && y < len(game) {
		if v, b := game[y-1][x].(empty); b {
			targets = append(targets, v)
		}
	}
	if x-1 < len(game[0]) && y < len(game) {
		if v, b := game[y][x-1].(empty); b {
			targets = append(targets, v)
		}
	}
	if x+1 < len(game[0]) && y < len(game) {
		if v, b := game[y][x+1].(empty); b {
			targets = append(targets, v)
		}
	}
	if y+1 < len(game[0]) && y < len(game) {
		if v, b := game[y+1][x].(empty); b {
			targets = append(targets, v)
		}
	}
	for _, step := range targets {
		b, path := astar(step, t, game)
		if b && len(path) < minLen {
			next = step.at()
			minLen = len(path)
		}
	}
	return next
}

func buildPath(cameFrom map[coordinate]coordinate, current coordinate) []empty {
	path := []empty{empty{current}}
	for {
		next, b := cameFrom[current]
		if !b {
			break
		}
		path = append(path, empty{next})
		current = next
	}
	forward := []empty{}
	for i := len(path) - 2; i >= 0; i-- {
		forward = append(forward, path[i])
	}
	return forward
}

func astar(startTail tail, targetTail tail, game [][]tail) (bool, []empty) {
	start := startTail.at()
	target := targetTail.at()
	closedSet := map[coordinate]bool{}
	openSet := map[coordinate]bool{}
	openSet[start] = true
	cameFrom := map[coordinate]coordinate{}
	gScore := map[coordinate]int{}
	gScore[start] = 0
	fScore := map[coordinate]int{}
	fScore[start] = dist(start, target)

	for len(openSet) > 0 {
		minWeight := math.MaxInt64
		var c coordinate
		for candidate := range openSet {
			weight := math.MaxInt64 - 1
			if w, b := fScore[candidate]; b {
				weight = w
			}
			if weight < minWeight {
				minWeight = weight
				c = candidate
			}
			if weight == minWeight && (candidate.y < c.y || (candidate.y == c.y && candidate.x < c.x)) {
				minWeight = weight
				c = candidate
			}
		}
		if c == target {
			return true, buildPath(cameFrom, c)
		}
		delete(openSet, c)
		closedSet[c] = true
		for _, neighbor := range []tail{game[c.y+1][c.x], game[c.y][c.x+1], game[c.y][c.x-1], game[c.y-1][c.x]} {
			if _, b := closedSet[neighbor.at()]; b {
				continue
			}
			if _, b := neighbor.(empty); !b {
				continue
			}
			tentativeWeight := gScore[c] + dist(c, neighbor.at())

			if _, b := openSet[neighbor.at()]; !b {
				openSet[neighbor.at()] = true
			} else if tentativeWeight >= gScore[neighbor.at()] {
				continue
			}
			cameFrom[neighbor.at()] = c
			gScore[neighbor.at()] = tentativeWeight
			fScore[neighbor.at()] = gScore[neighbor.at()] + dist(neighbor.at(), target)
		}
	}
	return false, nil
}

func inPath(needle tail, haystack []empty) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func deploy(p *player, game [][]tail) (bool, [][]tail) {
	enemies := findEnemies(*p, game)
	if len(enemies) < 1 {
		return false, game
	}
	targets := findTargets(game, enemies)
	targetFound, target := findClosest(*p, game, targets)
	if targetFound {
		nextStep := nextStep(*p, game, target)
		np := *p
		game[np.at().y][np.at().x] = empty{np.at()}

		if np.render() == 'E' {
			np = elf{nextStep, np.hpLeft()}
		} else {
			np = goblin{nextStep, np.hpLeft()}
		}
		game[nextStep.y][nextStep.x] = np
		*p = np
		return true, game
	}
	return false, game
}

func engage(p player, game [][]tail) (bool, player) {
	x := p.at().x
	y := p.at().y
	sorroundings := []tail{game[y-1][x], game[y][x-1], game[y][x+1], game[y+1][x]}
	enemyFound := false
	var weakestEnemy player
	life := math.MaxInt64
	for _, t := range sorroundings {
		if e, b := t.(player); b && p.enemy() == e.render() && e.hpLeft() > 0 {
			enemyFound = true
			if e.hpLeft() < life {
				weakestEnemy = e
				life = e.hpLeft()
			}
			if e.hpLeft() == life && (e.at().y < weakestEnemy.at().y || e.at().y == weakestEnemy.at().y && e.at().x < weakestEnemy.at().x) {
				weakestEnemy = e
				life = e.hpLeft()
			}
		}
	}
	return enemyFound, weakestEnemy
}

func attack(t *player, game [][]tail, x player) [][]tail {
	w := *t
	if w.hpLeft() < 1 {
		fmt.Println(w)
		fmt.Println(x)
		fmt.Println(game[w.at().y][w.at().x])
		renderGame(game, []empty{empty{coordinate{13, 21}}, empty{coordinate{1, 1}}})
		panic("Already dead")
	}
	if w.render() == 'E' {
		w = elf{w.at(), w.hpLeft() - 3}
	} else {
		w = goblin{w.at(), w.hpLeft() - (3 + bonus)}
	}
	*t = w
	if w.hpLeft() < 1 {
		game[w.at().y][w.at().x] = empty{w.at()}
	} else {
		game[w.at().y][w.at().x] = w
	}
	return game
}

func play(p *player, game [][]tail, players Players) (bool, [][]tail) {
	if canAttack, target := engage(*p, game); canAttack {
		var realTarget *player
		for _, pt := range players {
			casted := *pt
			if target.at() == casted.at() && casted.hpLeft() > 0 {
				realTarget = pt
			}
		}
		return true, attack(realTarget, game, target)
	}
	canMove := false
	canMove, game = deploy(p, game)

	canAttack := false
	if canAttack, target := engage(*p, game); canMove && canAttack {
		var realTarget *player
		for _, pt := range players {
			casted := *pt
			if target.at() == casted.at() && casted.hpLeft() > 0 {
				realTarget = pt
			}
		}
		return true, attack(realTarget, game, target)
	}
	return canMove || canAttack, game
}

func renderGame(game [][]tail, path []empty) {
	fmt.Printf("\033[0;0H")
	for y := 0; y < len(game); y++ {
		for x := 0; x < len(game[y]); x++ {
			if inPath(game[y][x], path) {
				fmt.Printf("\033[33;1m%c\033[0m", game[y][x].render())
			} else {
				fmt.Printf("%c", game[y][x].render())
			}
		}
		fmt.Println()
	}
}

func countElvesAlive(players []*player) int {
	count := 0
	for _, player := range players {
		p := *player
		if p.render() == 'E' && p.hpLeft() > 0 {
			count++
		}
	}
	return count
}

func battleSimulation(game [][]tail, players Players) (int, int, Players) {
	sort.Sort(players)
	c := 0
	for {
		shouldContinue := false
		for _, player := range players {
			tmp := *player
			if tmp.hpLeft() < 1 {
				shouldContinue = shouldContinue || false
				continue
			}
			canPlay := false
			canPlay, game = play(player, game, players)
			shouldContinue = shouldContinue || canPlay
		}
		sort.Sort(players)
		if !shouldContinue {
			tot := 0
			for _, player := range players {
				tmp := *player
				if tmp.hpLeft() > 0 {
					tot += tmp.hpLeft()
				}
			}
			return c, (c - 1) * tot, players
		}
		c++

	}
	return c, -1, players
}

func makeElvesWin(input []string) (int, int, Players) {
	game, players := readMap(input)
	totElves := countElvesAlive(players)
	rounds, result, units := battleSimulation(game, players)
	totSurvivors := countElvesAlive(units)
	fmt.Printf("Simulation %d ended with result %d in %d rounds, # survivors %d. Attacking with bonus %d\n", 1, result, rounds, totSurvivors, bonus)
	c := 1
	for totSurvivors < totElves {
		players = nil
		game, players = readMap(input)
		bonus++
		rounds, result, units = battleSimulation(game, players)
		totSurvivors = countElvesAlive(units)
		c++
		fmt.Printf("Simulation %d ended with result %d in %d rounds, # survivors %d. Attacking with bonus %d\n", c, result, rounds, totSurvivors, bonus)
	}
	return rounds, result, players
}

var bonus = 0

func main() {
	lines := readInput(bufio.NewScanner(os.Stdin))
	rounds, result, _ := makeElvesWin(lines)
	fmt.Printf("Simulation ended with result %d acchieved in %d rounds. Attacking with bonus %d\n", result, rounds, bonus)

	// tests()
}

func tests() {
	input := `#######   
			#.G...# 
			#...EG# 
			#.#.#G# 
			#..G#E# 
			#.....# 
			#######`
	rounds, result := runTest(input)
	if result != 27730 {
		fmt.Errorf("result for game \n %s \n incorrect, expected %d got %d", input, 27730, result)
		fmt.Println(rounds)
	}

	input = `#######
			#G..#E#
			#E#E.E#
			#G.##.#
			#...#E#
			#...E.#
			#######`
	rounds, result = runTest(input)
	if result != 36334 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 36334, result)
		fmt.Println(rounds)
	}
	input = `#######  
			#E..EG#  
			#.#G.E#  
			#E.##E#  
			#G..#.#  
			#..E#.#  
			#######`
	rounds, result = runTest(input)
	if result != 39514 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 39514, result)
		fmt.Println(rounds)
	}
	input = `####### 
			#E.G#.# 
			#.#G..# 
			#G.#.G# 
			#G..#.# 
			#...E.# 
			#######`
	rounds, result = runTest(input)
	if result != 27755 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 27755, result)
		fmt.Println(rounds)
	}
	input = `#######   
			#.E...#   
			#.#..G#   
			#.###.#  
			#E#G#G#   
			#...#G#   
			#######`
	rounds, result = runTest(input)
	if result != 28944 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 28944, result)
		fmt.Println(rounds)
	}
	input = `######### 
			#G......# 
			#.E.#...# 
			#..##..G# 
			#...##..# 
			#...#...# 
			#.G...G.# 
			#.....G.# 
			#########`
	rounds, result = runTest(input)
	if result != 18740 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 18740, result)
		fmt.Println(rounds)
	}

	input = `####### 
			#.G...# 
			#...EG# 
			#.#.#G# 
			#..G#E# 
			#.....# 
			#######`
	rounds, result, players := runTest2(input)
	if result != 4988 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 4988, result)
		fmt.Printf("Number of rounds + 1: %d\n", rounds)
		fmt.Printf("Attack needed: %d\n", bonus+3)
		for _, player := range players {
			fmt.Println(*player)
		}
	}
	bonus = 0
	input = `################################
	#G..#####G.#####################
	##.#####...#####################
	##.#######..####################
	#...#####.#.#.G...#.##...###...#
	##.######....#...G..#...####..##
	##....#....G.........E..####.###
	#####..#...G........G...##....##
	######.....G............#.....##
	######....G.............#....###
	#####..##.......E..##.#......###
	########.##...........##.....###
	####G.G.......#####..E###...####
	##.......G...#######..#####..###
	#........#..#########.###...####
	#.G..GG.###.#########.##...#####
	#...........#########......#####
	##..........#########..#.#######
	###G.G......#########....#######
	##...#.......#######.G...#######
	##.......G....#####.E...#.######
	###......E..G.E......E.....#####
	##.#................E.#...######
	#....#...................#######
	#....#E........E.##.#....#######
	#......###.#..#..##.#....#..####
	#...########..#..####....#..####
	#...########.#########......####
	#...########.###################
	############.###################
	#########....###################
	################################`
	rounds, result = runTest(input)
	if result != 206720 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 206720, result)
		fmt.Println(rounds)
	}
	bonus = 0
	input = `####### 
			#E..EG# 
			#.#G.E# 
			#E.##E# 
			#G..#.# 
			#..E#.# 
			#######`
	rounds, result, players = runTest2(input)
	if result != 31284 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 31284, result)
		fmt.Printf("Number of rounds + 1: %d\n", rounds)
		fmt.Printf("Attack needed: %d\n", bonus+3)
		for _, player := range players {
			fmt.Println(*player)
		}
	}
	bonus = 0
	input = `####### 
			#E.G#.# 
			#.#G..# 
			#G.#.G# 
			#G..#.# 
			#...E.# 
			#######`
	rounds, result, players = runTest2(input)
	if result != 3478 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 3478, result)
		fmt.Printf("Number of rounds + 1: %d\n", rounds)
		fmt.Printf("Attack needed: %d\n", bonus+3)
		for _, player := range players {
			fmt.Println(*player)
		}
	}
	bonus = 0
	input = `####### 
			#.E...# 
			#.#..G# 
			#.###.# 
			#E#G#G# 
			#...#G# 
			#######`
	rounds, result, players = runTest2(input)
	if result != 6474 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 6474, result)
		fmt.Printf("Number of rounds + 1: %d\n", rounds)
		fmt.Printf("Attack needed: %d\n", bonus+3)
		for _, player := range players {
			fmt.Println(*player)
		}
	}
	bonus = 0
	input = `######### 
			#G......# 
			#.E.#...# 
			#..##..G# 
			#...##..# 
			#...#...# 
			#.G...G.# 
			#.....G.# 
			#########`
	rounds, result, players = runTest2(input)
	if result != 1140 {
		fmt.Printf("result for game \n %s \n incorrect, expected %d got %d\n", input, 1140, result)
		fmt.Printf("Number of rounds + 1: %d\n", rounds)
		fmt.Printf("Attack needed: %d\n", bonus+3)
		for _, player := range players {
			fmt.Println(*player)
		}
	}

}

func runTest(input string) (int, int) {
	lines := readInput(bufio.NewScanner(strings.NewReader(input)))
	world, players := readMap(lines)
	rounds, result, _ := battleSimulation(world, players)
	return rounds, result
}

func runTest2(input string) (int, int, Players) {
	lines := readInput(bufio.NewScanner(strings.NewReader(input)))
	return makeElvesWin(lines)
}
