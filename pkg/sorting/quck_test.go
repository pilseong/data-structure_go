package sorting

import (
	"testing"
)

func TestQuickSort(t *testing.T) {

	arr := []int{20, 3, 32, 99, 53, 2, 5, 6, 8, 23, 21, 235, 10, 45, 33, int(^uint(0) >> 1)}

	QuickSort(arr, 0, len(arr)-1)

}
