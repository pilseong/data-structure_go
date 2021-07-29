package sorting

import (
	"log"
	"testing"
)

func TestMergeSort(t *testing.T) {
	arr := []int{3, 2, 1, 7, 30, 5, 24, 18, 9}

	var j, p int
	

	for p = 2; (p-1)*2 <= len(arr); p = p * 2 {
		for j = 0; j+p <= len(arr); j = j + p {
			log.Print(j, (j+j+p)/2, j+p, p)
			Merge(arr, j, (j+j+p)/2, j+p)
		}

		if len(arr) % p != 0 {
			log.Print("etc ", j, (j+j+p)/2, j+p, p)
			Merge(arr, j, (j+j+p)/2, len(arr))
		}
	}

	if p/2 < len(arr) {
		Merge(arr, 0, p/2, len(arr))
	}

	for _, d := range arr {
		log.Println(d)
	}
}
