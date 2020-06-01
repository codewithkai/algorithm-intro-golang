package arraylist

import "testing"

func TestArrayList(t *testing.T) {
	arrayList := ArrayList{}
	arrayList.Add(1)
	if arrayList.size != 1 {
		t.Error("array size should be 1")
	}
	expected := []int{1}
	if !sameOrder(arrayList.elements, expected) {
		t.Errorf("elements expected %v got %v", expected, arrayList.elements)
	}
	arrayList.Add(2)
	arrayList.Add(3)
	arrayList.Add(4)
	arrayList.Add(5)
	if arrayList.size != 5 {
		t.Error("array size should be 5")
	}
	expected = []int{1, 2, 3, 4, 5, 0, 0, 0}
	if !sameOrder(arrayList.elements, expected) {
		t.Errorf("elements expected %v got %v", expected, arrayList.elements)
	}
	arrayList.Remove(3)
	if arrayList.size != 4 {
		t.Error("array size should be 4")
	}
	expected = []int{1, 2, 4, 5, 0, 0, 0, 0}
	if !sameOrder(arrayList.elements, expected) {
		t.Errorf("elements expected %v got %v", expected, arrayList.elements)
	}

}

func sameOrder(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
