package main

import "crypto/sha256"

var EmptyNode = &Node{
	Hash: make([]byte, 32),
}

type Node struct {
	isLeaf bool
	Hash   []byte
	Parent *Node
	Left   *Node
	Right  *Node
}

func (node *Node) ComputeHash() []byte {
	if node.isLeaf {
		return node.Hash
	}

	hash := sha256.Sum256(append(node.Left.Hash, node.Right.Hash...))

	return hash[:]
}
