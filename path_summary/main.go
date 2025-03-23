package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) newnode(value int) *Node {
	newNode := &Node{Value: value}
	if t.root == nil {
		t.root = newNode
		return newNode
	}

	return newNode
}

func (t *Tree) addnode(ns [][]int) {
	var queueRoot []*Node
	var queueTmp []*Node
	cur := 0

	for i, inner := range ns {
		queueTmp = []*Node{}
		cur = 0
		for j, v := range inner {
			nn := t.newnode(v)
			queueTmp = append(queueTmp, nn)

			if i > 0 {
				q := queueRoot[cur]
				if q.Left == nil {
					q.Left = nn
				} else {
					q.Right = nn

					if j < i {
						cur++
						nq := queueRoot[cur]
						nq.Left = nn
					}
				}
			}
		}
		queueRoot = queueTmp
	}
}

func (t *Tree) summary() int {
	path := traversal(t.root)

	sum := 0
	for _, v := range path {
		sum += v
	}

	fmt.Println("path", path)
	return sum
}

var c []int

func traversal(node *Node) []int {
	if node == nil {
		return []int{}
	}
	c = append(c, node.Value)

	var left int
	if node.Left != nil {
		left = node.Left.Value
	}

	var right int
	if node.Right != nil {
		right = node.Right.Value
	}

	if left > right {
		traversal(node.Left)
	} else {
		traversal(node.Right)
	}

	return c
}

func main() {
	ns := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	tree := &Tree{}
	tree.addnode(ns)

	o := tree.summary()
	fmt.Printf("output = %d", o)
}
