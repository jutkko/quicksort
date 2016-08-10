package quicksort

import (
	"fmt"
	"math/rand"
	"reflect"
	builtinSort "sort"
	"testing"
)

func TestEmptyIntArray(t *testing.T) {
	intArray := []int{}
	Sort(intArray)
	if len(intArray) > 0 {
		t.Errorf("Failed to sort empty array")
	}
}

func TestOrderedIntArray(t *testing.T) {
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

func TestUnorderedIntArray(t *testing.T) {
	intArray := []int{5, 2, 3, 4, 1}
	Sort(intArray)
	fmt.Printf("%#v\n", intArray)
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

func TestLargeIntArray(t *testing.T) {
	count := 10000
	intArray := make([]int, 0, count)
	for i := 0; i < count; i++ {
		intArray = append(intArray, rand.Intn(10*count))
	}

	copyIntArray := make([]int, count, count)
	copy(copyIntArray, intArray)

	Sort(intArray)
	builtinSort.Ints(copyIntArray)
	if !reflect.DeepEqual(intArray, copyIntArray) {
		t.Errorf("Sort is still borked!")
	}
}
