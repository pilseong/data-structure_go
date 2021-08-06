package linkedlist

import (
	"log"
	"testing"

	"github.com/heops79/data-structure_go/pkg/model"
)


func Test_linkedlist(t *testing.T) {
	linkedList := LinkedList{
		Head: nil,
		Tail: nil,
	}

	linkedList.Append(&model.Entry{Key: "pilseong1", Value: 10})
	linkedList.Append(&model.Entry{Key: "pilseong2", Value: 20})
	linkedList.Append(&model.Entry{Key: "pilseong3", Value: 30})
	linkedList.Prepend(&model.Entry{Key: "pilseong4", Value: 5})
	linkedList.Prepend(&model.Entry{Key: "pilseong5", Value: 1})

	var _, n = linkedList.Search(5)
	log.Println(n)

	var deleted = linkedList.Remove(30)
	linkedList.Print()

	log.Printf("deleted: %d", deleted.Data.Value)
}
