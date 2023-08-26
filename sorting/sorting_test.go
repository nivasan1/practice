package sorting_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"leetcode.com/leetcode/sorting"
	"leetcode.com/leetcode/utils"
)

func TestInsertionSort(t *testing.T) {
	ints := []utils.OrdInt{utils.NewOrdInt(5), utils.NewOrdInt(3), utils.NewOrdInt(7)}
	sorting.InsertionSort[int](ints)
	assert.Equal(t, ints[0].Val(), 7)
}

func TestMergeSort(t *testing.T) {
	ints := []utils.OrdInt{utils.NewOrdInt(5), utils.NewOrdInt(3), utils.NewOrdInt(7), utils.NewOrdInt(13), utils.NewOrdInt(8)}
	sorting.MergeSort[int](ints)
	fmt.Println(ints)
}
