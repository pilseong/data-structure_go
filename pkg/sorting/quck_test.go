package sorting

import (
	"log"
	"math"
	"testing"
)

func TestQuickSort(t *testing.T) {

	arr := []int{20, 3, 32, 99, 53, 2, 5, 6, 8, 23, 21, 235, 10, 45, 33, math.MaxInt64}

	QuickSort(arr, 0, len(arr)-1)

	log.Println(arr)
}
