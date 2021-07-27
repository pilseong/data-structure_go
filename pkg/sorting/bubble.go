package sorting

func BubbleSort(arr []int) {
	var i, j int

	for i = 0; i < len(arr)-1; i++ {
		for j = i; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

}