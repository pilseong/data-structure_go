package sorting

func RadixSort(arr []int) {
	var beans [10]([]int)

	// flag is used for checking futher iteration needed.
	var flag = true

	// get the each digit for each iteration from the least important to the most
	for i := 1; flag; i++ {
		flag = false

		for _, v := range arr {
			child := (v % Power(10, i))
			parent := (Power(10, i) / 10)

			r := child / parent
			beans[r] = append(beans[r], v)

			// flag is set to true when all child is larger than parent
			if !flag && child > parent {
				flag = true
			}
		}

		// copy data from the bucket to the original only when the meaningful data is in the bucket
		if flag {
			arr = arr[:0]
			for index, s := range beans {
				for j := 0; j < len(s); j++ {
					arr = append(arr, s[j])
				}
				beans[index] = beans[index][:0]
			}
		}
	}
}

// Power is making power of times based on base
func Power(base, times int) int {

	var result = 1
	for i := 0; i < times; i++ {
		result = result * base
	}

	return result
}
