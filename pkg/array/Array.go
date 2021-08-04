package array

import "fmt"

// Array is using slice for base data-structure
type Array struct {
	arr []int
}

func (a *Array) Add(v int) {
	a.arr = append(a.arr, v)
}

func (a *Array) Display() {
	for _, v := range a.arr {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func (a *Array) Insert(i, v int) {

	// if the index is the same as length of the array then add the new value to the end of the array
	// it would be add value to the empty array
	if len(a.arr) == i {
		a.arr = append(a.arr, v)
		return
	}
	a.arr = append(a.arr[:i+1], a.arr[i:]...)
	a.arr[i] = v
}

func (a *Array) Delete(i int) {
	if i >= 0 && i < len(a.arr) {
		a.arr = append(a.arr[:i], a.arr[i+1:]...)
	}
}