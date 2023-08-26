package skip_list

import "leetcode.com/leetcode/utils"

// start with a regular sorted singly-linked list
type listNode[C comparable] struct {
	next *listNode[C]
	val  C
}

type List[C comparable] interface {
	Insert(C)
	Remove(C) (C, bool)
	Head() (C, bool)
	Get(find C) (C, bool)
	Tail() (C, bool)
}

type list[C comparable] struct {
	head *listNode[C]
	tail *listNode[C]
}

func NewList[C comparable]() List[C] {
	return &list[C]{}
}

func (l *list[C]) insertAt(val C, prev *listNode[C]) {
	// by default insert node in front of head
	n := &listNode[C]{
		val:  val,
		next: l.head,
	}
	// if the node is to be inserted after prev, update list
	if prev != nil {
		n.next = prev.next
		prev.next = n
	}
	// if the node was inserted before the head, update
	if n.next == l.head {
		l.head = n
	}
	// if there is no tail or the element was inserted in front of the tail update
	if l.tail == nil {
		l.tail = l.head.next
		return
	} else {
		if l.tail.next != nil {
			l.tail = l.tail.next
		}
	}
}

func (l *list[C]) Insert(val C) {
	// always insert at head
	l.insertAt(val, nil)
}

func (l *list[C]) get(pred func(*listNode[C]) bool) []*listNode[C] {
	var (
		node  = l.head
		nodes = make([]*listNode[C], 0)
	)
	// while there is a node to search
	for node != nil {
		// check if the predicate is met, if so, add the node in order
		if pred(node) {
			nodes = append(nodes, node)
		}
		node = node.next
	}
	return nodes
}

func (l *list[C]) Get(find C) (C, bool) {
	var val C
	node := l.get(func(n *listNode[C]) bool {
		return n.val == find
	})
	if len(node) == 0 {
		return val, false
	}
	return node[0].val, true
}

func (l *list[C]) Remove(remove C) (C, bool) {
	var val C
	// find the element, and its predecessor
	nodes := l.get(func(n *listNode[C]) bool {
		if n != nil {
			if n.val == remove {
				return true
			}
			if n.next != nil {
				return n.next.val == remove
			}
		}
		return false
	})
	if len(nodes) == 0 {
		return val, false
	}
	// if there is only one node, the node we are searching for is at the head
	if len(nodes) == 1 {
		val = nodes[0].val
		l.head = nodes[0].next
		return val, true
	}
	val = nodes[0].val
	// skip the second node given (node matching value)
	nodes[0].next = nodes[1].next
	l.tail = nil
	if nodes[0].next == nil && l.head != nodes[0] {
		l.tail = nodes[0]
	}
	return val, true
}

func (l *list[C]) Head() (C, bool) {
	var val C
	if l.head == nil {
		return val, false
	}
	return l.head.val, true
}

func (l *list[C]) Tail() (C, bool) {
	var val C
	if l.tail == nil {
		return val, false
	}
	return l.tail.val, true
}

type SortedList[C comparable, O utils.Ord[C]] interface {
	List[C]
	SortInsert(O)
	SortGet(O) (C, bool)
}

var _ SortedList[int, utils.Ord[int]] = (*sortedList[int, utils.Ord[int]])(nil)

func NewSortedList[C comparable, O utils.Ord[C]]() SortedList[C, O] {
	return &sortedList[C, O]{
		NewList[C]().(*list[C]),
	}
}

type sortedList[C comparable, O utils.Ord[C]] struct {
	*list[C]
}

func (s *sortedList[C, O]) SortInsert(val O) {
	// identify previous element
	var prev *listNode[C]
	nodes := s.get(func(ln *listNode[C]) bool {
		return !val.Less(ln.val)
	})
	if len(nodes) != 0 {
		prev = nodes[len(nodes)-1]
	}
	s.insertAt(val.Val(), prev)
}

func (s *sortedList[C, O]) SortGet(val O) (C, bool) {
	return s.Get(val.Val())
}

// no-op
func (s *sortedList[C, O]) Insert(val C) {}
