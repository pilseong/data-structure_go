package sorting

func partition(arr []int, l, h int) int {

	var i, j int
	i = l + 1
	j = h

	for {
		for {
			if arr[i] > arr[l] {
				break
			}
			i++
		}

		for {
			if arr[j] <= arr[l] {
				break
			}
			j--
		}

		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			arr[l], arr[j] = arr[j], arr[l]
			return j
		}
	}
}

func QuickSort(arr []int, l, h int) {

	if l == h {
		return
	}
	// log.Println(arr, l, h)
	pivot := partition(arr, l, h)
	// log.Println(arr, l, pivot)
	// log.Println()

	QuickSort(arr, l, pivot)
	QuickSort(arr, pivot+1, h)

}
