package linkedlist

import (
	"log"

	"github.com/heops79/data-structure_go/pkg/model"
)


type LinkedList struct {
	Head *model.Node
	Tail *model.Node
	Size int
}

func (l *LinkedList) Search(data int) (pre *model.Node, found *model.Node) {
	cur := l.Head
	var prev *model.Node

	for cur != nil {
		if cur.Data.Value == data {
			return prev, cur
		}
		prev = cur
		cur = cur.Next
	}
	return nil, nil
}

func (l *LinkedList) Remove(data int) *model.Node {
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

func (l *LinkedList) Append(data *model.Entry) {
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

func (l *LinkedList) Prepend(data *model.Entry) {
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

func createNode(data *model.Entry) *model.Node {
	newNode := &model.Node{
		Data: data,
		Next: nil,
	}

	return newNode
}

func (l *LinkedList) Print() {
	var cur = l.Head
	for cur != nil {
		log.Printf("%v, next: %v", cur.Data, cur.Next)
		cur = cur.Next
	}
}
