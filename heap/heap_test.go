package heap_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"leetcode.com/leetcode/heap"
	"leetcode.com/leetcode/utils"
)

func TestMaxHeap(t *testing.T) {
	// create new max-heap
	heap := heap.NewMaxHeap(func(a, b int) bool {
		return a >= b
	},
		func() int {
			return utils.MIN_INT
		},
	)

	elements := []int{23, 17, 14, 6, 13, 5, 7, 10, 1}

	for _, element := range elements {
		heap.Insert(element)
	}

	slices.Sort(elements)
	slices.Reverse(elements)

	for _, element := range elements {
		front := heap.PopFront()
		assert.Equal(t, front, element)
	}
}

func TestMaxHeapify(t *testing.T) {
	// give random array
	elements := []int{1, 2, 3, 4, 5, 6, 7, 8}

	h := heap.MaxHeapify(elements,
		func(a, b int) bool {
			return a >= b
		},
		func() int {
			return utils.MIN_INT
		},
	)

	for i := 8; i >= 1; i-- {
		next := h.PopFront()
		assert.Equal(t, i, next)
	}
}
