package searching

import (
	"sort"
)

func Search(target int, arr []int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func BinarySearch(target int, arr []int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if target < arr[mid] {
			high = mid - 1
		} else if target > arr[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

//use sort package
func BinarySearch2(target int, arr []int) int {
	i := sort.Search(len(arr), func(i int) bool { return arr[i] >= target })
	if i < len(arr) && arr[i] == target {
		return i
	}
	return -1
}
