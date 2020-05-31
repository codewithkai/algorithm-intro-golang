package sorting

import "math/rand"


func QuickSort(arr []int) []int{
	if len(arr) <= 1{
		return arr
	}
	pivotValue := arr[0]
	var small, big []int
	for i:=1; i<len(arr); i++{
		if arr[i] <= pivotValue{
			small = append(small, arr[i])
		}else{
			big = append(big, arr[i])
		}
	}
	//quicksort(small) + pivotValue + quicksort(big)
	return append(append(QuickSort(small), pivotValue), QuickSort(big)...)
}



// in place
func QuickSortInPlace(arr []int){
	if len(arr) <= 1{
		return
	}
	pivotIndex := rand.Intn(len(arr))

	//swap pivot and 0
	arr[0], arr[pivotIndex] = arr[pivotIndex], arr[0]
	last := 0
	for i:=1; i<len(arr);i++{
		if arr[i] < arr[0]{
			//swap i and last
			last++
			arr[i], arr[last] = arr[last], arr[i]
		}
	}
	arr[last], arr[0] = arr[0], arr[last]
	QuickSortInPlace(arr[:last])
	QuickSortInPlace(arr[last+1:])
}
