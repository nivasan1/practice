package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"leetcode.com/leetcode/utils"
)

func TestLog2(t *testing.T) {
	assert.Equal(t, utils.Log2(4), 2)
	assert.Equal(t, utils.Log2(3), 1)
	assert.Equal(t, utils.Log2(8), 3)
	assert.Equal(t, utils.Log2(11), 2)
}

func TestIsPow2(t *testing.T) {
	assert.False(t, utils.IsPow2(3))
	assert.True(t, utils.IsPow2(16))
	assert.False(t, utils.IsPow2(9))
}
