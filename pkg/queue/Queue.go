package queue

import "github.com/heops79/data-structure_go/pkg/model"

type Queue struct {
	Head *model.Node
	Tail *model.Node
	Size int
}

func (q *Queue) Enqueue(data int) bool {
	var node = &model.Node{
		Data: data,
		Next: nil,
	}

	if (q.Size == 0) {
		q.Head = node
		q.Tail = node
	} else {
		node.Next = q.Head
		q.Head = node
	}

	q.Size++

	return true
}

