package main

import "crypto/sha256"

type MerkleTree struct {
	root *Node
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
		leafs = append(leafs, leafs[len(leafs)-1])
	}

	return &MerkleTree{root: build(leafs)}
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

func (tree *MerkleTree) Hash() []byte {
	return tree.root.Hash
}
