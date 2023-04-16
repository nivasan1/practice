package skip_list_test

import (
	"github.com/stretchr/testify/assert"
	sk "leetcode.com/leetcode/skip_list"
	"leetcode.com/leetcode/utils"
	"testing"
)

func TestListInsertion(t *testing.T) {
	// create a new list
	list := sk.NewList[int]()
	// insert new element
	list.Insert(1)
	// check that head is equal to tail
	head, ok := list.Head()
	assert.True(t, ok)
	assert.Equal(t, head, 1)
	_, ok = list.Tail()
	assert.False(t, ok)
	// insert another list element
	list.Insert(2)
	head, ok = list.Head()
	assert.True(t, ok)
	assert.Equal(t, head, 2)
	tail, ok := list.Tail()
	assert.True(t, ok)
	assert.Equal(t, tail, 1)
}

func TestListMultipleInsertRemoval(t *testing.T) {
	// create new list
	list := sk.NewList[int]()
	// insert two elements
	list.Insert(1)
	list.Insert(2)
	head, ok := list.Head()
	assert.True(t, ok)
	assert.Equal(t, head, 2)
	tail, ok := list.Tail()
	assert.True(t, ok)
	assert.Equal(t, tail, 1)
	// insert another element
	list.Insert(3)
	// remove elements (3 -> 2 -> 1)
	list.Remove(2)
	head, _ = list.Head()
	assert.Equal(t, head, 3)
	assert.Equal(t, tail, 1)
	// attempt to get 2
	_, ok = list.Get(2)
	assert.False(t, ok)
	// remove the last element
	_, ok = list.Remove(1)
	assert.True(t, ok)
	_, ok = list.Tail()
	assert.False(t, ok)
	val, ok := list.Head()
	assert.True(t, ok)
	assert.Equal(t, val, 3)
	// remove head
	_, ok = list.Remove(3)
	assert.True(t, ok)
	_, ok = list.Head()
	assert.False(t, ok)
}

func TestInsertRemoveInsert(t *testing.T) {
	list := sk.NewList[int]()
	list.Insert(1)
	list.Insert(2)
	val, ok := list.Head()
	assert.Equal(t, val, 2)
	assert.True(t, ok)
	val, ok = list.Tail()
	assert.Equal(t, val, 1)
	assert.True(t, ok)
	list.Remove(1)
	list.Remove(2)
	_, ok = list.Head()
	assert.False(t, ok)
	_, ok = list.Tail()
	assert.False(t, ok)
	// insert an element
	list.Insert(3)
	val, ok = list.Head()
	assert.True(t, ok)
	assert.Equal(t, val, 3)
}

// sorted list test
func TestInsertion(t *testing.T) {
	list := sk.NewSortedList[int, utils.OrdInt]()
	// insert elements
	list.SortInsert(utils.NewOrdInt(1))
	list.SortInsert(utils.NewOrdInt(4))
	// check head / tail
	head, ok := list.Head()
	assert.True(t, ok)
	assert.Equal(t, 1, head)
	tail, ok := list.Tail()
	assert.True(t, ok)
	assert.Equal(t, 4, tail)
	// insert 3, shld be in between
	list.SortInsert(utils.NewOrdInt(3))
	tail, _ = list.Tail()
	assert.Equal(t, 4, tail)
	_, ok = list.Get(3)
	assert.True(t, ok)
	// remove the last element
	_, ok = list.Remove(4)
	assert.True(t, ok)
	// check that tail is updated to 3
	tail, ok = list.Tail()
	assert.True(t, ok)
	assert.Equal(t, tail, 3)
}
