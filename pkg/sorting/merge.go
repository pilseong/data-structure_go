package sorting

import "log"

func Merge(arr []int, l, mid, h int) {

	log.Println(l, mid, h)
	var i, j, k int
	var res []int


	for j=mid; i < mid && j < h; k++ {
		if arr[i] > arr[j] {
			res[k] = arr[j]
			j++
		} else {
			res[k] = arr[i]
			i++
		}
	}

	for i < mid {
		res[k] = arr[i]
		i++
		k++
	}

	for j < h {
		res[k] = arr[j]
		j++
		k++
	}

}