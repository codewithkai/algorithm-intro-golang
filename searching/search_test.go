package searching

import (
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

type SearchFunc func(int, []int) int

func TestSearch(t *testing.T) {
	arr := []int{1,3,5,7,9}
	data := []struct {
		target  int
		expectedIndex int
	}{
		{1, 0},
		{3, 1},
		{8,-1},
	}
	functions := []SearchFunc{Search, BinarySearch, BinarySearch2}
	for _, f := range functions{
		for _, v := range data {
			functionName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
			foundIndex := f(v.target, arr)
			if foundIndex != v.expectedIndex {
				t.Errorf("search func %s target %d index expected %d got %d",
					functionName,  v.target, v.expectedIndex, foundIndex)
			}
		}
	}
}

func TestLargeInput(t *testing.T){
	inputSize := 10000000
	arr := make([]int, inputSize)
	for i:=0; i< inputSize; i++{
		arr[i] = rand.Intn(inputSize)
	}
	functions := []SearchFunc{Search, BinarySearch, BinarySearch2}
	for _, f := range functions{
		functionName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		now := time.Now()
		f(rand.Intn(inputSize), arr)
		duration := time.Since(now)
		t.Logf("func %s duration %v", functionName,duration)
	}
}

