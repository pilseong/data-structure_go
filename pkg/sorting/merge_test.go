package sorting

import (
	"fmt"
	"testing"
)


func TestMergeSort(t *testing.T) {
	arr := []int{1, 3, 5, 2, 4, 6, 8}

	Merge(arr, 0, len(arr)/2, len(arr))

	for _, d := range arr {
		fmt.Println(d)
	}
}