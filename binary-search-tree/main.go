package main

import "fmt"

func printTree(tree *BinarySearchTree) {
	fmt.Println("tree()")

	printRoot(tree.root)
}

func printRoot(root *Node) {
	if root.left != nil {
		printRoot(root.left)
	}

	fmt.Printf("  node(%d)\n", root.data)

	if root.right != nil {
		printRoot(root.right)
	}
}

func main() {
	tree := NewBinarySearchTree()
	tree.Insert(10)
	tree.Insert(8)
	tree.Insert(20)
	tree.Insert(18)
	tree.Insert(2)

	values := []int{2, 8, 10, 18, 20}

	for _, value := range values {
		if tree.Contains(value) {
			fmt.Printf("ok: tree contains: %d\n", value)
		} else {
			fmt.Printf("fail: tree should contain 18 but reports it doesn't: %d\n", value)
		}
	}

	printTree(tree)
}
