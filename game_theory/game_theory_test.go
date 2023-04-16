package game_theory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"leetcode.com/leetcode/game_theory"
)

func TestMinimax(t *testing.T) {
	assert.Equal(t, game_theory.Minimax([]int{3, 5, 2, 9, 12, 5, 23, 23}, true), 12)
}
