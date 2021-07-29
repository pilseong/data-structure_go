package sorting

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{3, 2, 1, 7, 30, 5, 24} //, 18, 9, 11}

	MergeSort(arr)

	for _, d := range arr {
		log.Println(d)
	}
}
