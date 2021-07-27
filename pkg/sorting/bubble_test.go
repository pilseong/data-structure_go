package sorting

import (
	"fmt"
	"testing"
)


func TestBubbleSort(t *testing.T) {
	arr := []int{3, 2, 5, 13, 34, 1, 4}

	BubbleSort(arr)

	for _, d := range arr {
		fmt.Println(d)
	}
}