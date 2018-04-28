package main

import "fmt"

const (
	RED   string = "red"
	BLACK string = "black"
)

var nilNode = &Node{color: BLACK}

type Node struct {
	value int
	color string
	left  *Node
	right *Node
}

func (n *Node) withRed() {
	n.color = RED
}

func (n *Node) withBlack() {
	n.color = BLACK
}

func (n *Node) isRed() bool {
	return n.color == RED
}

func (n *Node) isBlack() bool {
	return n.color == BLACK
}

func (n *Node) isNilNode() bool {
	return n == nilNode
}

func (n *Node) valueToStrng() string {
	if n.value == 0 {
		return ""
	}

	if n.isRed() {
		return fmt.Sprintf("\033[1;31m%d\033[0;39m", n.value)
	} else {
		return fmt.Sprintf("%d", n.value)
	}
}

func newRootNode(v int) *Node {
	return &Node{value: v, color: BLACK, left: nilNode, right: nilNode}
}

func newNode(v int) *Node {
	return &Node{value: v, color: RED, left: nilNode, right: nilNode}
}

type Path struct {
	list []*Node
}

func (p *Path) pop() *Node {
	if p.noList() {
		return nil
	}

	n := p.list[len(p.list)-1]
	p.list = p.list[:len(p.list)-1]
	return n
}

func (p *Path) noList() bool {
	return len(p.list) == 0
}
