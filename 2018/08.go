package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stack []node

type node struct {
	children  []node
	nchildren int
	meta      int
	start     int
	end       int
}

func (s stack) Push(v node) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, node) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack) Peek() node {
	l := len(s)
	return s[l-1]
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	stringLicense, _ := scanner.ReadString('\n')
	tokens := strings.Split(stringLicense, " ")
	license := []int{}
	for _, v := range tokens {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		license = append(license, i)
	}

	tree := buildTree(license)
	_, n := tree.Pop()
	result := one_08(license, n)
	fmt.Printf("result = %d\n", result)
	fmt.Println("---------")
	value := two_08(license, n)
	fmt.Printf("result = %d\n", value)
}

func sum(a []int) int {
	t := 0
	for _, f := range a {
		t += f
	}
	return t
}

func one_08(input []int, tree node) int {
	if tree.nchildren == 0 {
		return sum(input[tree.end-tree.meta+1 : tree.end+1])
	} else {
		result := sum(input[tree.end-tree.meta+1 : tree.end+1])
		for _, child := range tree.children {
			result += one_08(input, child)
		}
		return result
	}
}

func two_08(input []int, tree node) int {
	if tree.nchildren == 0 {
		return sum(input[tree.end-tree.meta+1 : tree.end+1])
	} else {
		value := 0
		indexes := input[tree.end-tree.meta+1 : tree.end+1]
		for _, i := range indexes {
			if i <= len(tree.children) && i > 0 {
				value += two_08(input, tree.children[i-1])
			}
		}
		return value
	}
}

func buildTree(input []int) stack {
	tree := make(stack, 0)
	i := 0
	for i < len(input) {
		if input[i] > 0 {
			tree = tree.Push(node{[]node{}, input[i], input[i+1], i, -1})
			i += 2
		} else {
			mama := node{}
			tree, mama = tree.Pop()
			child := node{[]node{}, 0, input[i+1], i, i + 2 + input[i+1] - 1}
			mama.children = append(mama.children, child)
			tree = tree.Push(mama)
			p := child.end
			closableMama := tree.Peek()
			grandma := node{}
			for closableMama.nchildren == len(closableMama.children) {
				tree, mama = tree.Pop()
				mama.end = p + mama.meta
				p += mama.meta
				if len(tree) == 0 {
					tree = tree.Push(mama)
					break
				} else {
					tree, grandma = tree.Pop()
					grandma.children = append(grandma.children, mama)
					tree = tree.Push(grandma)
					closableMama = tree[len(tree)-1]
				}
			}
			i = p + 1
		}
	}

	return tree
}
