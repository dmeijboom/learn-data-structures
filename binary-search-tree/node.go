package main

type Node struct {
	left  *Node
	right *Node
	data  int
}

func NewNode(data int) *Node {
	return &Node{data: data}
}
