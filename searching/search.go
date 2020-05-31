package searching

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

