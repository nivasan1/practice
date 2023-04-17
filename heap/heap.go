package heap

import (
	"leetcode.com/leetcode/utils"
)

// abstraction around nodes in heap
type HeapNode interface {
	Parent() HeapNode
	RightChild() HeapNode
	LeftChild() HeapNode
	index() int
}

type heapNodeImpl struct {
	i int
}

func (h heapNodeImpl) Parent() HeapNode {
	return heapNodeImpl{
		i: h.i >> 1,
	}
}

func (h heapNodeImpl) RightChild() HeapNode {
	return heapNodeImpl{
		i: h.i<<1 + 1,
	}
}

func (h heapNodeImpl) LeftChild() HeapNode {
	return heapNodeImpl{
		i: h.i << 1,
	}
}

func (h heapNodeImpl) index() int {
	return h.i
}

type Heap[C comparable, O utils.Ord[C]] interface {
	Insert(O)
	PopFront() C
	PopBack() C
	GetAt(HeapNode) O
}

type heapImpl[C comparable, O utils.Ord[C]] struct {
	// an array of ordered elements
	arr      []O
	length   int
	heapSize int
}

func NewHeap[C comparable, O utils.Ord[C]](length int) Heap[C, O] {
	return nil
}

func (h *heapImpl[C, O]) Insert(val O) {
	// push element at back of heap
	if h.heapSize >= h.length {
		return
	}
	h.heapSize++
	h.arr[h.heapSize] = val
	// ensure max-heap invariant is maintained, create HeapNodes
	var heapNode HeapNode
	heapNode = heapNodeImpl{i: h.heapSize}
	curVal := h.GetAt(heapNode)
	// if the current value is less than the parent
	for !curVal.Less(h.GetAt(heapNode.Parent()).Val()) {
		// curVal will now be placed at the index of its parent, and its parent will be placed at curVal's current index
		h.arr[heapNode.index()], h.arr[heapNode.Parent().index()] = h.arr[heapNode.Parent().index()], h.arr[heapNode.index()]
		heapNode = heapNode.Parent()
	}
}

func (h *heapImpl[C, O]) GetAt(n HeapNode) O {
	var val O
	idx := n.index()
	if idx <= h.heapSize {
		val = h.arr[idx]
	}
	return val
}

func (h *heapImpl[C, O]) PopFront() C {
	var val C
	// if there are elements in the heap
	if h.heapSize >= 0 {

	}
}
