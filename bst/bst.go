// Package bst implements a binary search tree.
package bst

type Node struct {
	left  *Node
	right *Node
	value int
}

func NewBST(rootVal int) Node {
	n := Node{nil, nil, rootVal}
	return n
}

func (n *Node) Exists(item int) bool {

}

func (n *Node) Remove(item int) (Node, bool) {

}

func (n *Node) Add(item int) (Node, bool) {
}

func (n *Node) IsBST() bool {

}
