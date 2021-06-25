package main

type BinarySearchTree struct {
	root *Node
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (tree *BinarySearchTree) Insert(value int) {
	if tree.root == nil {
		tree.root = NewNode(value)
		return
	}

	current := tree.root

	for current != nil {
		if value < current.data {
			if current.left == nil {
				current.left = NewNode(value)
				break
			}

			current = current.left
			continue
		}

		if current.right == nil {
			current.right = NewNode(value)
			break
		}

		current = current.right
	}
}

func (tree *BinarySearchTree) Contains(value int) bool {
	if tree.root == nil {
		return false
	}

	current := tree.root

	for current != nil {
		if current.data == value {
			return true
		}

		if value < current.data {
			current = current.left
			continue
		}

		current = current.right
	}

	return false
}
