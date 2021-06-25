package main

type Node struct {
	isLeaf bool
	Hash   []byte
	Parent *Node
	Left   *Node
	Right  *Node
}
