package game_theory

import "leetcode.com/leetcode/utils"

// given a set of terminal scores, and whether the final turn is made by the minimizer / maximizer
func Minimax(scores []int, maximizer bool) int {
	// leafs must be power of 2 (perfect binary tree)
	if !utils.IsPow2(len(scores)) {
		return -1
	}
	// base-case
	if len(scores) == 1 {
		return scores[0]
	}
	// minimaxed_scores is the scores chosen by the min/max imizer
	minimaxedScores := make([]int, len(scores)/2)
	// every entry in the array represents a leaf
	j := 0
	var fn func(int, int) int
	for i := 0; i < len(scores); i += 2 {
		if maximizer {
			fn = max
		} else {
			fn = min
		}
		minimaxedScores[j] = fn(scores[i], scores[i+1])
		j++
	}
	return Minimax(minimaxedScores, !maximizer)
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
