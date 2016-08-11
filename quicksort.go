package quicksort

import "math/rand"

func Sort(array []int) {
	if len(array) == 0 {
		return
	}

	sort(array, 0, len(array)-1)
}

func sort(array []int, s, e int) {
	if s >= e {
		return
	}

	var pivot, g, u int
	g = s
	u = s
	pivotIndex := rand.Intn(e-s) + s
	array[pivotIndex], array[e] = array[e], array[pivotIndex]
	pivot = array[e]

	for ; u < e; u++ {
		if array[u] <= pivot {
			array[u], array[g] = array[g], array[u]
			g++
		}
	}

	array[g], array[e] = array[e], array[g]

	sort(array, s, g-1)
	sort(array, g+1, e)
}
