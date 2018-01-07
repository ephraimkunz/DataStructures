// Package bst implements a binary search tree.
package bst

import (
	"strconv"
	"strings"
)

type BST struct {
	root  *node
	count int
}

type node struct {
	value int
	left  *node
	right *node
}

// NewBST creates and returns a new, empty binary search tree
func NewBST() *BST {
	return &BST{}
}

// Exists returns true if item in bst, else false.
func (bst *BST) Exists(item int) bool {
	node, _ := recursiveFind(item, bst.root)
	return node != nil
}

// Remove removes item from bst and returns true if succeeded,
// false if item not in bst.
func (bst *BST) Remove(item int) bool {
	child, parent := recursiveFind(item, bst.root)

	if child == nil { // Not in tree
		return false
	}

	if parent == nil { // Removing root
		if bst.root.left != nil {
			right := bst.root.right
			bst.root = bst.root.left

			// Put right side in correct location
			insertPos := bst.root
			for ; insertPos.right != nil; insertPos = insertPos.right {
			}
			insertPos.right = right
		} else if bst.root.right != nil {
			left := bst.root.left
			bst.root = bst.root.right

			// Put left side in correct location
			insertPos := bst.root
			for ; insertPos.left != nil; insertPos = insertPos.left {
			}
			insertPos.left = left
		} else {
			bst.root = nil
		}
		bst.count--
		return true
	}
	//return false

	if parent.left == child {
		if child.left != nil {
			right := child.right
			parent.left = child.left

			// Put right side in correct location
			insertPos := parent.left
			for ; insertPos.right != nil; insertPos = insertPos.right {
			}
			insertPos.right = right
		} else if child.right != nil {
			left := child.left
			parent.left = child.right

			// Put left side in correct location
			insertPos := parent.left
			for ; insertPos.left != nil; insertPos = insertPos.left {
			}
			insertPos.left = left
		} else {
			parent.left = nil
		}
		bst.count--
		return true
	} else {
		if child.left != nil {
			right := child.right
			parent.right = child.left

			// Put right side in correct location
			insertPos := parent.right
			for ; insertPos.right != nil; insertPos = insertPos.right {
			}
			insertPos.right = right
		} else if child.right != nil {
			left := child.left
			parent.right = child.right

			// Put left side in correct location
			insertPos := parent.right
			for ; insertPos.left != nil; insertPos = insertPos.left {
			}
			insertPos.left = left
		} else {
			parent.right = nil
		}
		bst.count--
		return true
	}
}

// Searchest tree with root node for item. If item is found,
// child points to it and parent points to its parent. If the found
// node is the root of the entire tree, the parent is nil. If the item is not found,
// child is nil and parent holds the parent of where the child should go.
func recursiveFind(item int, root *node) (child, parent *node) {
	if root == nil {
		return nil, nil
	}

	if item == root.value { // Found item is the root of the entire tree
		return root, nil
	}

	if item < root.value {
		if root.left == nil {
			return nil, root
		} else if item == root.left.value {
			return root.left, root
		} else {
			return recursiveFind(item, root.left)
		}
	} else {
		if root.right == nil {
			return nil, root
		} else if item == root.right.value {
			return root.right, root
		} else {
			return recursiveFind(item, root.right)
		}
	}
}

// Insert inserts the item correctly into the bst, but does not tree balancing.
// No duplicates are allowed, so Insert returns false if item is a duplicate.
func (bst *BST) Insert(item int) bool {
	// First item
	if bst.root == nil {
		bst.root = &node{item, nil, nil}
		bst.count++
		return true
	}

	child, parent := recursiveFind(item, bst.root)
	if child != nil {
		return false
	}

	if item < parent.value {
		parent.left = &node{item, nil, nil}
	} else {
		parent.right = &node{item, nil, nil}
	}
	bst.count++
	return true
}

// For checking state in testing
func (bst *BST) statePrint() string {
	result := statePrintRecursive(bst.root)

	// Strip off possible leading and trailing whitespace
	return strings.TrimSpace(result)
}

func statePrintRecursive(root *node) string {
	if root == nil { // Base case
		return ""
	}

	return statePrintRecursive(root.left) + " " + strconv.Itoa(root.value) + statePrintRecursive(root.right)
}
