package main

import (
	"log"

	"github.com/heops79/data-structure_go/pkg/linkedlist"
)

func main() {

	linkedList := linkedlist.LinkedList{
		Head: nil,
		Tail: nil,
	}

	linkedList.Append(10)
	linkedList.Append(20)
	linkedList.Append(30)
	linkedList.Prepend(5)
	linkedList.Prepend(1)

	var _, n = linkedList.Search(5)
	log.Println(n)

	var deleted = linkedList.Remove(30)
	linkedList.Print()

	log.Printf("deleted: %d", deleted.Data)

}
