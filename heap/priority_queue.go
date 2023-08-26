package heap

type Stringer interface {
	String() string
}

type PriorityQueue[A Stringer] interface {
	Insert(a A, priority int)
	PopFront() A
}

type priorityQueueElement[A Stringer] struct {
	value    A
	priority int
}

type priorityQueueImpl[A Stringer] struct {
	// cache value.String() -> idx in heap
	valueCache map[string]int

	// underlying heap of pqes
	heap *maxHeap[*priorityQueueElement[A]]
}

func (pq *priorityQueueImpl[A]) Insert(a A, priority int) {
	// check if value exists in cache
	idx, ok := pq.valueCache[a.String()]
	if ok {
		// change priority of element in heap
		oldPriority := pq.heap.elements[idx].priority
		pq.heap.elements[idx].priority = priority

		// if priority is lowered, only need to check that heap rooted at idx is valid
		if oldPriority > priority {
			pq.heap.maxHeapIfy(pq.heap.elements[idx:], idx)
		}
		// if new priority is increased, check that parent of idx is valid heap
		if priority > oldPriority {
			pq.heap.maxHeapIfy(pq.heap.elements[parent(idx):], parent(idx))
		}

		return
	}

	// create the element, and insert
	pqe := &priorityQueueElement[A]{
		value:    a,
		priority: priority,
	}

	pq.heap.Insert(pqe)
}

func (pq *priorityQueueImpl[A]) PopFront() A {
	val := pq.heap.PopFront()
	delete(pq.valueCache, val.value.String())
	return val.value
}
