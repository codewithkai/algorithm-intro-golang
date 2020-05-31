package sorting

import "testing"

type Entry struct{
	input []int
	expected []int
}



func TestQuickSort(t *testing.T){
	data := getTestData()
	for _, v := range data{
		output := QuickSort(v.input)
		if !sameOrder(output, v.expected){
			t.Errorf("expected %v got %v", v.expected, v.input)
		}
	}
}

func TestQuickSortInPlace(t *testing.T){
	data := getTestData()
	for _, v := range data{
		QuickSortInPlace(v.input)
		if !sameOrder(v.input, v.expected){
			t.Errorf("expected %v got %v", v.expected, v.input)
		}
	}
}

func getTestData() []Entry{
	return []Entry{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1,7,3,8,9}, []int{1,3,7,8,9}},
		{[]int{1,2,3,4,5}, []int{1,2,3,4,5}},
	}
}

func sameOrder(arr1 []int, arr2 []int) bool{
	if len(arr1)  != len(arr2){
		return false
	}
	for i:=0; i< len(arr1); i++{
		if arr1[i] != arr2[i]{
			return false
		}
	}
	return true
}
