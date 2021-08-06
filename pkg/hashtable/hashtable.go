package hashtable

import (
	"log"

	"github.com/heops79/data-structure_go/pkg/linkedlist"
	"github.com/heops79/data-structure_go/pkg/model"
)



type Hashtable struct {
	Arr []*linkedlist.LinkedList
	BucketSize int
}

func (h *Hashtable) hash(key string) int {
	s := len(key)
	for i:=0; i<len(key); i++ {
		s = (s + int(key[i])) * i 
	}
	
	return s % h.BucketSize
}

func (h *Hashtable) Put(key string, value int) {

	i := h.hash(key)

	log.Println("hash index", i)

	if h.Arr[i] == nil {
		h.Arr[i] = new(linkedlist.LinkedList)
	} 
	h.Arr[i].Append(&model.Entry{key, value})
}

func (h *Hashtable) Get(key string) int {
	// h.Arr[key].Search(key)
	return 0
}
