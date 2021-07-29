package sorting

import "log"

// Merge two different sections inside an array
// one is from [l, mid), another is from [mid, h)
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

	// slice is totally different from array
	if i < mid {
		res = append(res, arr[i:mid]...)
		i++
		k++
	}

	if j < h {
		res = append(res, arr[j:h]...)
		j++
		k++
	}

	log.Println("log", res)

	for i = 0; i < len(res); i++ {
		arr[l+i] = res[i]
	}

	// log.Println((arr))
}

func MergeSort(arr []int) {
	var j, p int

	for p = 2; p <= len(arr); p = p * 2 {

		for j = 0; j+p <= len(arr); j = j + p {
			log.Print(j, (j+j+p)/2, j+p, p)
			Merge(arr, j, (j+j+p)/2, j+p)
		}

		// if the size of array is not divided by power of 2
		// get the last slice, merge them
		if len(arr)%p != 0 {

			log.Print("l ", j, j+(len(arr)-j-(p/2)), len(arr), p)
			Merge(arr, j, j+(len(arr)-j-(p/2)), len(arr))
		}
	}

	// last merge for the remaining elements
	if p/2 < len(arr) {
		log.Print("e ", 0, p/2, len(arr), p)
		Merge(arr, 0, p/2, len(arr))
	}
}
