// A simple implementation of the binary-search tree data structure.

package main

import (
	"fmt"
)

// A node in the BST.
type Node struct {
	val   int
	left  *Node
	right *Node
}

// A BST.
type BST struct {
	root *Node
}

// Factory function to create new BSTs.
func newNode(val int, left, right *Node) *Node {
	return &Node{val: val, left: left, right: right}
}

// Do an in-order walk of the tree rooted at "root".
func (root *Node) inOrder(f func(*Node)) {
	if root.left != nil {
		root.left.inOrder(f)
	}
	f(root)
	if root.right != nil {
		root.right.inOrder(f)
	}
}

func printNode(node *Node) {
	fmt.Println(node.val)
}

// Factory function to create new BSTs.
func newBST() *BST {
	return &BST{root: nil}
}

// Insert an integer into a BST.
func (b *BST) insert(i int) {
	if b.root == nil {
		b.root = newNode(i, nil, nil)
		return
	}
	b.root.insert(i)
}

// Insert an integer into a BST represented by a "root" node.
func (root *Node) insert(i int) {
	if root.val < i {
		if root.right == nil {
			root.right = newNode(i, nil, nil)
		} else {
			root.right.insert(i)
		}
	} else if root.val > i {
		if root.left == nil {
			root.left = newNode(i, nil, nil)
		} else {
			root.left.insert(i)
		}
	} else { // Already present, we don't insert duplicates.
		return
	}
}

func main() {
	bst := newBST()
	vals := []int{3, 9, -2, 48, 6, 0, 15, -17}
	for _, v := range vals {
		bst.insert(v)
	}
	bst.root.inOrder(printNode)
}
