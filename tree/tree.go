package tree

// Tree has one root and branches and leafs
type Tree *Node

// Node is a part of tree
type Node struct {
	Descendants map[rune]*Node
}
