package model

type Node struct {
	Data int
	Next *Node
}

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}