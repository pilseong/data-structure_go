package linkedlist

import "log"

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *LinkedList) Search(data int) (pre *Node, found *Node) {
	cur := l.Head
	var prev *Node

	for cur != nil {
		if cur.Data == data {
			return prev, cur
		}
		prev = cur
		cur = cur.Next
	}
	return nil, nil
}

func (l *LinkedList) Remove(data int) *Node {
	var prev, found = l.Search(data)

	if found == nil {
		return nil
	}

	// found the target in the head
	if l.Head == found {
		l.Head = l.Head.Next
		// found the target in the tail
	} else if l.Tail == found {
		l.Tail = prev
		prev.Next = found.Next
		// found the target in the middle
	} else {
		prev.Next = found.Next
	}

	// set Head and Tail to nil when there is only one item in the list
	if l.Size == 1 {
		l.Head = nil
		l.Tail = nil
	}

	// set the node not in the list
	found.Next = nil
	l.Size--

	return found
}

func (l *LinkedList) Append(data int) {
	var newNode = createNode(data)

	if l.Head == nil && l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}

func (l *LinkedList) Prepend(data int) {
	var newNode = createNode(data)

	if l.Head == nil && l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Size++
}

func createNode(data int) *Node {
	newNode := Node{
		Data: data,
		Next: nil,
	}

	return &newNode
}

func (l *LinkedList) Print() {
	var cur = l.Head
	for cur != nil {
		log.Printf("%v, next: %v", cur, cur.Next)
		cur = cur.Next
	}
}
