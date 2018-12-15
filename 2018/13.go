package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

type Position struct {
	x int
	y int
}
type Step struct {
	up    Position
	down  Position
	right Position
	left  Position
	t     string
}
type Cart struct {
	current   Position
	direction rune
	nextTurn  int
	active    bool
}

type CartSort []Cart

func (c CartSort) Len() int      { return len(c) }
func (c CartSort) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CartSort) Less(i, j int) bool {
	return c[i].current.y < c[j].current.y || c[i].current.y == c[j].current.y && c[i].current.x < c[j].current.x
}

type Direction struct {
	left  rune
	right rune
}

var relative = map[rune]Direction{}
var null = Position{-1, -1}

func getStep(p Position, c rune, course map[Position]Step) Step {
	up := Position{p.x, p.y - 1}
	down := Position{p.x, p.y + 1}
	right := Position{p.x + 1, p.y}
	left := Position{p.x - 1, p.y}
	if c == '-' {
		return Step{up: null, down: null, right: right, left: left, t: "n"}
	}
	if c == '|' {
		return Step{up: up, down: down, right: null, left: null, t: "n"}
	}
	if c == '+' {
		return Step{up: up, down: down, right: right, left: left, t: "i"}
	}
	if c == '/' {
		if prev, b := course[left]; b && prev.left != null && prev.right != null {
			return Step{up: up, down: null, right: null, left: left, t: "n"}
		}
		return Step{up: null, down: down, right: right, left: null, t: "n"}
	}
	if c == '\\' {
		if prev, b := course[left]; b && prev.left != null && prev.right != null {
			return Step{up: null, down: down, right: null, left: left, t: "n"}
		}
		return Step{up: up, down: null, right: right, left: null, t: "n"}
	}
	return Step{up: up, down: down, right: right, left: left, t: "n"}
}

func getCart(p Position, c rune) Cart {
	return Cart{p, c, 0, true}
}

func cleanCourse(course map[Position]Step) {
	for p := range course {
		if _, there := course[course[p].up]; !there {
			step := course[p]
			step.up = null
			course[p] = step
		}
		if _, there := course[course[p].down]; !there {
			step := course[p]
			step.down = null
			course[p] = step
		}
		if _, there := course[course[p].right]; !there {
			step := course[p]
			step.right = null
			course[p] = step
		}
		if _, there := course[course[p].left]; !there {
			step := course[p]
			step.left = null
			course[p] = step
		}
	}
}

func printCourse(course map[Position]Step, x int, y int, carts []Cart) {
	canvas := [][]rune{}
	for i := 0; i <= y; i++ {
		canvas = append(canvas, []rune{})
		for j := 0; j <= x; j++ {
			canvas[i] = append(canvas[i], ' ')
		}
	}
	for p := range course {
		if course[p].t == "i" {
			canvas[p.y][p.x] = '+'
		} else {
			canvas[p.y][p.x] = '.'
		}
	}

	for _, cart := range carts {
		if !cart.active {
			continue
		}
		a := canvas[cart.current.y][cart.current.x]
		if a == '^' || a == '<' || a == '>' || a == 'v' {
			canvas[cart.current.y][cart.current.x] = 'X'
		} else {
			canvas[cart.current.y][cart.current.x] = cart.direction
		}
	}

	fmt.Printf("\033[0;0H")
	for _, row := range canvas {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	maxx := 0
	y := 0
	course := map[Position]Step{}
	carts := []Cart{}
	relative['^'] = Direction{left: '<', right: '>'}
	relative['v'] = Direction{left: '>', right: '<'}
	relative['<'] = Direction{left: 'v', right: '^'}
	relative['>'] = Direction{left: '^', right: 'v'}
	for scanner.Scan() {
		line := scanner.Text()
		x := 0
		for _, c := range line {
			if c == ' ' {
				x++
				continue
			}
			pos := Position{x, y}
			step := getStep(pos, c, course)
			if c == '^' || c == 'v' || c == '<' || c == '>' {
				carts = append(carts, getCart(pos, c))
			}
			if (Step{}) == step {
				log.Printf("Error reading char %c", c)
			} else {
				course[pos] = step
			}
			x++
		}
		if x > maxx {
			maxx = x
		}
		y++
	}

	cleanCourse(course)
	active := len(carts)
	for active > 1 {
		drive(course, carts)
		active = countRemoved(carts)
		sort.Sort(CartSort(carts))
	}
	for _, cart := range carts {
		if cart.active {
			fmt.Printf("last cart standing at [%d,%d]\n", cart.current.x, cart.current.y)
		}
	}

}

func countRemoved(carts []Cart) int {
	count := 0
	for _, cart := range carts {
		if cart.active {
			count++
		}
	}
	return count
}

func drive(course map[Position]Step, carts []Cart) {
	for i := 0; i < len(carts); i++ {
		carts[i] = move(carts[i], course)
		for a := 0; a < len(carts); a++ {
			if a != i && carts[a].active && carts[i].active && carts[a].current == carts[i].current {
				carts[a].active = false
				carts[i].active = false
				fmt.Printf("Crash at [%d,%d]\n", carts[a].current.x, carts[a].current.y)
				break
			}
		}
	}
}

func move(cart Cart, course map[Position]Step) Cart {
	if course[cart.current].t == "i" {
		if cart.nextTurn == 0 {
			cart.direction = relative[cart.direction].left
		} else if cart.nextTurn == 2 {
			cart.direction = relative[cart.direction].right
		} else {
		}
		cart.nextTurn = (cart.nextTurn + 1) % 3
	}
	d := cart.direction
	if d == '>' {
		if course[cart.current].right != null {
			cart.current = course[cart.current].right
		} else if course[cart.current].up != null {
			cart.current = course[cart.current].up
			cart.direction = '^'
		} else {
			cart.current = course[cart.current].down
			cart.direction = 'v'
		}
	} else if d == '<' {
		if course[cart.current].left != null {
			cart.current = course[cart.current].left
		} else if course[cart.current].up != null {
			cart.current = course[cart.current].up
			cart.direction = '^'
		} else {
			cart.current = course[cart.current].down
			cart.direction = 'v'
		}
	} else if d == '^' {
		if course[cart.current].up != null {
			cart.current = course[cart.current].up
		} else if course[cart.current].left != null {
			cart.current = course[cart.current].left
			cart.direction = '<'
		} else {
			cart.current = course[cart.current].right
			cart.direction = '>'
		}
	} else if d == 'v' {
		if course[cart.current].down != null {
			cart.current = course[cart.current].down
		} else if course[cart.current].left != null {
			cart.current = course[cart.current].left
			cart.direction = '<'
		} else {
			cart.current = course[cart.current].right
			cart.direction = '>'
		}
	}

	return cart
}
