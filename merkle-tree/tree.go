package main

import (
	"bytes"
	"crypto/sha256"
)

type MerkleTree struct {
	root  *Node
	leafs []*Node
}

func build(nodes []*Node) *Node {
	if len(nodes) == 1 {
		return nodes[0]
	}

	parentNodes := []*Node{}

	for i := 0; i < len(nodes); i += 2 {
		left := nodes[i]
		right := nodes[i+1]

		data := append(left.Hash, right.Hash...)
		hash := sha256.Sum256(data)

		parent := &Node{
			Left:  left,
			Right: right,
			Hash:  hash[:],
		}

		left.Parent = parent
		right.Parent = parent

		parentNodes = append(parentNodes, parent)
	}

	return build(parentNodes)
}

func NewMerkleTree(contents [][]byte) *MerkleTree {
	leafs := make([]*Node, len(contents))

	for i, content := range contents {
		hash := sha256.Sum256(content)

		leafs[i] = &Node{
			isLeaf: true,
			Hash:   hash[:],
		}
	}

	if len(leafs)%2 == 1 {
		leafs = append(leafs, EmptyNode)
	}

	return &MerkleTree{
		root:  build(leafs),
		leafs: leafs,
	}
}

func (tree *MerkleTree) Hash() []byte {
	return tree.root.Hash
}

func (tree *MerkleTree) Verify(content []byte) bool {
	hash := sha256.Sum256(content)

	for _, leaf := range tree.leafs {
		if !bytes.Equal(leaf.Hash, hash[:]) {
			continue
		}

		parent := leaf.Parent

		for parent != nil {
			leftHash := parent.Left.ComputeHash()
			rightHash := parent.Right.ComputeHash()
			hash := sha256.Sum256(append(leftHash, rightHash...))

			if !bytes.Equal(parent.Hash, hash[:]) {
				return false
			}

			parent = parent.Parent
		}

		return true
	}

	return false
}
