package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	if i <= bst.data {
		if bst.left == nil {
			bst.left = &BinarySearchTree{data: i}
		} else {
			bst.left.Insert(i)
		}
	} else {
		if bst.right == nil {
			bst.right = &BinarySearchTree{data: i}
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	var result []int

	var inorder func(node *BinarySearchTree)
	inorder = func(node *BinarySearchTree) {
		if node == nil {
			return
		}
		inorder(node.left)
		result = append(result, node.data)
		inorder(node.right)
	}

	inorder(bst)
	return result
}
