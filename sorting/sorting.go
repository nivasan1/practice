package sorting

import (
	"fmt"

	"leetcode.com/leetcode/utils"
)

func InsertionSort[C comparable, O utils.Ord[C]](arr []O) error {
	for i := range arr {
		if i+1 > len(arr) {
			return fmt.Errorf("length of array: %d exceeded by index %d", len(arr), i+1)
		}
		// sort array including i + 1
		insertSorted[C](arr[:i+1])
	}
	return nil
}

func insertSorted[C comparable, O utils.Ord[C]](arr []O) {
	// start from back and attempt to sort
	for i := len(arr) - 1; i > 0; i-- {
		// if the next element in the sequence is less than val
		if arr[i-1].LessOrd(arr[i]) {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		} else {
			break
		}
	}
	return
}

func MergeSort[C comparable, O utils.Ord[C]](arr []O) {
	mergeSort[C](arr, 0, len(arr)-1)
}

func mergeSort[C comparable, O utils.Ord[C]](arr []O, start, finish int) {
	if start < finish {
		// determine mid point
		mid := (finish + start) / 2
		// recursively sort first half
		mergeSort[C](arr, start, mid)
		// sort second half
		mergeSort[C](arr, mid+1, finish)
		// join both halves
		merge[C](arr, start, mid, finish)
	}
}

// for arr[p+1].Less(arr[p]), i.e arr shld be sorted in descending order
// we have that arr[p:q] in sorted order, and arr[q+1:r] in sorted order, want arr[p:r] to be in sorted order
func merge[C comparable, O utils.Ord[C]](arr []O, p, q, r int) {
	var x O
	L := make([]O, (q-p)+2)
	R := make([]O, r-q+1)
	// copy arrays
	copy(L, arr[p:q+1])
	copy(R, arr[q+1:r+1])
	// set sentinel values
	L[len(L)-1] = x.Max().(O)
	R[len(R)-1] = x.Max().(O)
	// update arr[p:r]
	i, j := 0, 0
	for k := p; k <= r; k++ {
		// update values
		var val O
		if L[i].LessOrd(R[j]) {
			val = L[i]
			i++
		} else {
			val = R[j]
			j++
		}
		arr[k] = val
	}
}
