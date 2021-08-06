package model

type Node struct {
	Data *Entry
	Next *Node
}

type Entry struct {
	Key   string
	Value int
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}