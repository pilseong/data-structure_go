package sorting

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {

	arr := []int{3, 2, 5, 13, 34, 1, 4}

	SelectionSort(arr)

	for _, d := range arr {
		fmt.Println(d)
	}

}