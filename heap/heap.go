package heap

type Heap[A any] interface {
	Insert(a A)
	PopFront() A
}

type maxHeap[A any] struct {
	// elements will be ordered array
	elements []A
	// return whether a is greater than b
	greater func(a, b A) bool
	// func min-val
	minVal func() A
}

func NewMaxHeap[A any](greater func(a, b A) bool, minVal func() A) Heap[A] {
	return &maxHeap[A]{
		greater: greater,
		minVal:  minVal,
	}
}

func MaxHeapify[A any](elements []A, greater func(a, b A) bool, minVal func() A) Heap[A] {
	// first create the max-heap
	heap := &maxHeap[A]{
		greater:  greater,
		minVal:   minVal,
		elements: elements,
	}

	len := len(elements)
	// start from the len >> 1 index (i.e children extend past the array, and max-heapify)
	for idx := len >> 1; idx >= 0; idx-- {
		heap.maxHeapIfy(elements, idx)
	}

	return heap
}

func Heapsort[A any](elements []A, greater func(a, b A) bool, minVal func() A) []A {
	// first construct heap
	heap := MaxHeapify(elements, greater, minVal).(*maxHeap[A])

	for i := len(elements); i >= 1; i-- {
		// max-heapify
		heap.maxHeapIfy(elements[:i], 0)
		// switch
		elements[0], elements[i-1] = elements[i-1], elements[0]
	}

	return elements
}

func (h *maxHeap[A]) Insert(a A) {
	// if the heap is empty
	if h.isEmpty() {
		h.elements = append(h.elements, a)
		return
	}

	// append element to end of array
	h.elements = append(h.elements, a)

	// now start at back, and continuously switch until heap property holds
	idx := len(h.elements) - 1
	for !h.greater(h.elements[parent(idx)], h.elements[idx]) {
		// swap
		h.elements[idx], h.elements[parent(idx)] = h.elements[parent(idx)], h.elements[idx]

		// update idx of new element
		idx = parent(idx)
	}
	return
}

func (h *maxHeap[A]) PopFront() A {
	var ret A
	if h.isEmpty() {
		return ret
	}
	ret = h.elements[0]
	// set root to min-int, and max-heapify
	h.elements[0] = h.minVal()

	// max-heapify
	h.maxHeapIfy(h.elements, 0)
	return ret
}

func (h *maxHeap[A]) maxHeapIfy(elements []A, i int) {
	if i >= len(elements) {
		return
	}

	largest := i

	if left(i) < len(elements) && h.greater(elements[left(i)], elements[largest]) {
		largest = left(i)
	}

	if right(i) < len(elements) && h.greater(elements[right(i)], elements[largest]) {
		largest = right(i)
	}

	// sub-tree rooted at elements[i] is currently a max-heap
	if largest == i {
		return
	}

	// swap elements[i] with largest
	elements[i], elements[largest] = elements[largest], elements[i]

	// continue
	h.maxHeapIfy(h.elements, largest)
}

func (h *maxHeap[A]) isEmpty() bool {
	return len(h.elements) == 0
}

func parent(i int) int {
	return i >> 1
}

func left(i int) int {
	return i << 1
}

func right(i int) int {
	return (i << 1) + 1
}
