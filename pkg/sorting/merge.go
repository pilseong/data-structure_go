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
