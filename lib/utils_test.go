package lib

import (
	"testing"
)

func TestSplice(t *testing.T) {
	var arr = []int{1, 5, 6, 7}

	arr = remove(arr, 2)

	if !equal(arr, []int{1, 5, 7}) {
		t.Errorf("%v", arr)
	}

	arr = unshift(arr, 1234)

	if !equal(arr, []int{1234, 1, 5, 7}) {
		t.Errorf("%v", arr)
	}
}
