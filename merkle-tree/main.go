package main

import (
	"fmt"
	"strings"
)

func printTree(tree *MerkleTree) {
	printNode(tree.root, 0)
}

func printNode(node *Node, level int) {
	fmt.Printf("%s%x\n", strings.Repeat("  ", level), node.Hash)

	if node.isLeaf {
		fmt.Printf("%sdata\n", strings.Repeat("  ", level+1))

		return
	}

	if node.Left != nil {
		printNode(node.Left, level + 1)
	}

	if node.Right != nil {
		printNode(node.Right, level+1)
	}
}

func main() {
	data := [][]byte{
		[]byte("item1"),
		[]byte("item2"),
		[]byte("item3"),
		[]byte("item4"),
	}

	tree := NewMerkleTree(data)

	printTree(tree)

	if tree.Verify([]byte("item3")) {
		fmt.Println("ok: verification succeeded for item3")
	} else {
		fmt.Println("fail: verification failed to item3")
	}

	if tree.Verify([]byte("item5")) {
		fmt.Println("fail: verification succeeded for item5")
	} else {
		fmt.Println("ok: verification failed to item5")
	}
}
