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

	linkedList.Append(&model.Entry{"pilseong1", 10})
	linkedList.Append(&model.Entry{"pilseong2", 20})
	linkedList.Append(&model.Entry{"pilseong3", 30})
	linkedList.Prepend(&model.Entry{"pilseong4", 5})
	linkedList.Prepend(&model.Entry{"pilseong5", 1})

	var _, n = linkedList.Search(5)
	log.Println(n)

	var deleted = linkedList.Remove(30)
	linkedList.Print()

	log.Printf("deleted: %d", deleted.Data.Value)
}
