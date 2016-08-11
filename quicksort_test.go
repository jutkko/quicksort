package quicksort

import (
	"math/rand"
	"reflect"
	builtinSort "sort"
	"testing"
)

const COUNT = 1000000

func TestEmptyArray(t *testing.T) {
	intArray := []int{}
	Sort(intArray)
	if len(intArray) > 0 {
		t.Errorf("Failed to sort empty array")
	}
}

func TestOrderedArray(t *testing.T) {
	intArray := []int{1, 2, 3, 4, 5}
	Sort(intArray)
	var prev, curr int
	for index, i := range intArray {
		if index > 0 {
			prev = intArray[index-1]
			curr = i
			if prev > curr {
				t.Errorf("At index %d: %d > %d", index-1, prev, curr)
			}
		}
	}
}

func TestUnorderedArray(t *testing.T) {
	intArray := []int{5, 2, 3, 4, 1}
	Sort(intArray)
	var prev, curr int
	for index, i := range intArray {
		if index > 0 {
			prev = intArray[index-1]
			curr = i
			if prev > curr {
				t.Errorf("At index %d: %d > %d", index-1, prev, curr)
			}
		}
	}
}

func TestLargeArray(t *testing.T) {
	intArray := make([]int, 0, COUNT)
	for i := 0; i < COUNT; i++ {
		intArray = append(intArray, rand.Intn(10*COUNT))
	}

	copyIntArray := make([]int, COUNT, COUNT)
	copy(copyIntArray, intArray)

	Sort(intArray)
	builtinSort.Ints(copyIntArray)
	if !reflect.DeepEqual(intArray, copyIntArray) {
		t.Errorf("Sort is still borked!")
	}
}

func BenchmarkLargeArray(b *testing.B) {
	intArray := make([]int, 0, COUNT)
	for i := 0; i < COUNT; i++ {
		intArray = append(intArray, rand.Intn(10*COUNT))
	}

	b.ResetTimer()
	Sort(intArray)
}

func BenchmarkLargeArrayBuiltinSort(b *testing.B) {
	intArray := make([]int, 0, COUNT)
	for i := 0; i < COUNT; i++ {
		intArray = append(intArray, rand.Intn(10*COUNT))
	}

	b.ResetTimer()
	builtinSort.Ints(intArray)
}

func BenchmarkLargeOrderedArray(b *testing.B) {
	intArray := make([]int, 0, COUNT)
	for i := 0; i < COUNT; i++ {
		intArray = append(intArray, i)
	}

	b.ResetTimer()
	Sort(intArray)
}

func BenchmarkLargeOrderedArrayBuiltinSort(b *testing.B) {
	intArray := make([]int, 0, COUNT)
	for i := 0; i < COUNT; i++ {
		intArray = append(intArray, i)
	}

	b.ResetTimer()
	builtinSort.Ints(intArray)
}
