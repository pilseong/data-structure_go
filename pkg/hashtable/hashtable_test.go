package hashtable

import (
	"testing"

	"github.com/heops79/data-structure_go/pkg/linkedlist"
)

func Test_hashtable(t *testing.T) {

	var h1 = Hashtable{
		Arr: make([]*linkedlist.LinkedList, 10),
		BucketSize: 10,
	}


	h1.Put("pilsonga", 100)
	
	h1.Arr[5].Print()
}