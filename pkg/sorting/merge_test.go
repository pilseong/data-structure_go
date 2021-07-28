package sorting

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{1, 3, 5, 7, 2, 4, 6, 8}

	var j, p int
	p = 1

	for p*2 <= len(arr) {
		p = p * 2

		for j = 0; j+p <= len(arr); j = j + p {
			log.Println(j, (j+j+p)/2, j+p)
			Merge(arr, j, (j+j+p)/2, j+p)
		}
	}

	Merge(arr, 0, len(arr)/2, len(arr))

	for _, d := range arr {
		log.Println(d)
	}
}
