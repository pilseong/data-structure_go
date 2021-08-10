package sorting

import (
	"log"
	"testing"
)

func TestRadixSort(t *testing.T) {
	arr := []int{20, 3, 32, 99, 53, 2, 5, 6, 8, 23, 21, 235, 10, 45, 33}

	log.Println("radix original", arr)

	RadixSort(arr)

	log.Println("radix result", arr)
}
