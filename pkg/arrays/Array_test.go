package arrays

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	var arr Array

	arr.Add(1)
	arr.Add(2)
	arr.Add(3)

	if len(arr.arr) != 3 {
		t.Errorf("Expected 3 but got %v", len(arr.arr))
	}
	log.Println(arr.arr)
}

func TestDisplay(t *testing.T) {
	var arr Array

	arr.Add(1)
	arr.Add(2)
	arr.Add(3)

	arr.Display()
}

func TestInsert(t *testing.T) {
	var arr Array

	arr.Add(1)
	arr.Add(3)
	arr.Add(4)

	arr.Display()

	arr.Insert(1, 2)

	if len(arr.arr) != 4 || arr.arr[1] != 2 {
		t.Errorf("Expected size 4 and the value 2 but got size of %v, value %v", len(arr.arr), arr.arr[1])
	}

	arr.Display()
}

func TestDelete(t *testing.T) {
	var arr Array

	arr.Add(1)
	arr.Add(2)
	arr.Add(3)

	arr.Delete(1)

	if len(arr.arr) != 2 {
		t.Errorf("Expected 2 but got %v", len(arr.arr))
	}
	arr.Display()
}
