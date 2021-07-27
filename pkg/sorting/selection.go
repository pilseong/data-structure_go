package sorting

func SelectionSort(arr []int) {

	var i, j, k int

	for ; i < len(arr)-1; i++ {
		k = i
		for j = i + 1; j < len(arr); j++ {
			if arr[k] > arr[j] {
				k = j
			}
		}
		arr[i], arr[k] = arr[k], arr[i]
	}
}