package lib

import "time"

func CurrentTimeMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func unshift(slice []int, v int) []int {
	var tmp = make([]int, len(slice)+1)
	tmp[0] = v
	copy(tmp[1:], slice)
	return tmp
}

func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
