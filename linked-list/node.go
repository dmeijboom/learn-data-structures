package main

type Node struct {
	next *Node
	data int
}

func NewNode(data int) *Node {
	return &Node{data: data}
}