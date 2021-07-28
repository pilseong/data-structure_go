package sorting

import "log"

func Merge(arr []int, l, mid, h int) {

	var i, j, k int
	var res []int

	i = l
	j = mid

	for ; i < mid && j < h; k++ {
		if arr[i] > arr[j] {
			res = append(res, arr[j])
			j++
		} else {
			res = append(res, arr[i])
			i++
		}
	}

	if i < mid {
		res = append(res, arr[i:]...)
		i++
		k++
	}

	if j < h {
		res = append(res, arr[j:]...)
		j++
		k++
	}

	copy(arr, res)
	log.Println((arr))
}
