package searching

import (
	"math/rand"
	"testing"
	"time"
)

type Entry struct {
	target  int
	expectedIndex int
}

func TestSearch(t *testing.T) {
	arr, data := getTestData()
	for _, v := range data {
		foundIndex := Search(v.target, arr)
		if foundIndex != v.expectedIndex {
			t.Errorf("target %d index expected %d got %d",
				v.target, v.expectedIndex, foundIndex)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	arr, data := getTestData()
	for _, v := range data {
		foundIndex := BinarySearch(v.target, arr)
		if foundIndex != v.expectedIndex {
			t.Errorf("target %d index expected %d got %d",
				v.target, v.expectedIndex, foundIndex)
		}
	}
}

func getTestData() ([]int, []Entry){
	arr := []int{1,3,5,7,9}
	entries := []Entry{
		{1, 0},
		{3, 1},
		{8,-1},
	}
	return arr, entries
}

func TestLargeInput(t *testing.T){
	inputSize := 10000000
	arr := make([]int, inputSize)
	for i:=0; i< inputSize; i++{
		arr[i] = rand.Intn(inputSize)
	}

	//Search
	now := time.Now()
	Search(rand.Intn(inputSize), arr)
	duration := time.Since(now)
	t.Logf("Search duration %v",duration)


	//Binary Search
	now = time.Now()
	BinarySearch(rand.Intn(inputSize), arr)
	duration = time.Since(now)
	t.Logf("BinarySearch duration %v",duration)
}

