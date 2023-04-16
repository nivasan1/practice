package prob_test

import (
	"github.com/stretchr/testify/assert"
	solution "leetcode.com/leetcode"
	"testing"
)

func Test1(t *testing.T) {
	ans := solution.ProductExceptSelf([]int{1, 2, 3})
	assert.Equal(t, ans, []int{6, 3, 2})
	ans = solution.ProductExceptSelf([]int{1, 2, 3, 4})
	assert.Equal(t, ans, []int{24, 12, 8, 6})
	ans = solution.ProductExceptSelf([]int{-1, 1, 0, -3, 3})
	assert.Equal(t, ans, []int{0, 0, 9, 0, 0})
}

func Test3(t *testing.T) {
	ans := solution.FindMin([]int{1, 2, 3})
	assert.Equal(t, ans, 1)
	ans = solution.FindMin([]int{3, 1, 2})
	assert.Equal(t, ans, 1)
	ans = solution.FindMin([]int{2, 3, 1})
	assert.Equal(t, ans, 1)
	ans = solution.FindMin([]int{2, 3, 4, 5, 1})
	assert.Equal(t, ans, 1)
	ans = solution.FindMin([]int{12, 13, 3, 7, 9})
	assert.Equal(t, ans, 3)
	ans = solution.FindMin([]int{3, 4, 5, 1, 2})
	assert.Equal(t, 1, ans)
	ans = solution.FindMin([]int{4, 5, 6, 7, 0, 1, 2})
	assert.Equal(t, ans, 0)
}

func Test4(t *testing.T) {
	ans := solution.Search([]int{2, 3, 4, 5, 6, 7, 1}, 4)
	assert.Equal(t, ans, 2)
	ans = solution.Search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(t, ans, 4)
	ans = solution.Search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	assert.Equal(t, ans, -1)
	ans = solution.Search([]int{11, 12, 13, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
	assert.Equal(t, ans, 12)
	ans = solution.Search([]int{11, 12, 13, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 6)
	assert.Equal(t, ans, 8)
	ans = solution.Search([]int{11, 12, 13, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 14)
	assert.Equal(t, ans, -1)
	ans = solution.Search([]int{1, 3}, 3)
	assert.Equal(t, ans, 1)
	ans = solution.Search([]int{1, 3}, 4)
	assert.Equal(t, ans, -1)
	ans = solution.Search([]int{5, 1, 3}, 1)
}

func Test5(t *testing.T) {
	ans := solution.ClimbStairs(3)
	assert.Equal(t, ans, 3)
}

func Test6(t *testing.T) {
	ans := solution.CoinChange([]int{1, 2, 5}, 11)
	assert.Equal(t, ans, 3)
	ans = solution.CoinChange([]int{2}, 3)
	assert.Equal(t, ans, -1)
	ans = solution.CoinChange([]int{1, 2, 5}, 10)
	assert.Equal(t, ans, 2)
	ans = solution.CoinChange([]int{1}, 2)
	assert.Equal(t, ans, 2)
}

func Test7(t *testing.T) {
	ans := solution.CoinChange2([]int{1, 2, 5}, 11)
	assert.Equal(t, ans, 3)
	ans = solution.CoinChange2([]int{2}, 3)
	assert.Equal(t, ans, -1)
	ans = solution.CoinChange2([]int{1, 2, 5}, 10)
	assert.Equal(t, ans, 2)
	ans = solution.CoinChange2([]int{1}, 2)
	assert.Equal(t, ans, 2)
	ans = solution.CoinChange2([]int{1, 2, 5}, 100)
	assert.Equal(t, ans, 20)
}

func Test8(t *testing.T) {
	ans := solution.LengthOfLIS([]int{1, 2, 3})
	assert.Equal(t, ans, 3)
	ans = solution.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	assert.Equal(t, ans, 4)
	ans = solution.LengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	assert.Equal(t, ans, 4)
	ans = solution.LengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	assert.Equal(t, ans, 1)
	ans = solution.LengthOfLIS([]int{4, 10, 4, 3, 8, 9})
	assert.Equal(t, ans, 3)
	ans = solution.LengthOfLIS([]int{1, 2, 2, 2})
	assert.Equal(t, ans, 2)
}

func Test9(t *testing.T) {
	ans := solution.LongestCommonSubsequence("abc", "def")
	assert.Equal(t, ans, 0)
	ans = solution.LongestCommonSubsequence("abcde", "ace")
	assert.Equal(t, ans, 3)
	ans = solution.LongestCommonSubsequence("abc", "abd")
	assert.Equal(t, ans, 3)
	ans = solution.LongestCommonSubsequence("abc", "bac")
	assert.Equal(t, ans, 2)
}

func Test10(t *testing.T) {
	ans := solution.TaskSchedulerII([]int{1, 2, 1, 2, 3, 1}, 3)
	assert.Equal(t, (int)(ans), 9)
	ans = solution.TaskSchedulerII([]int{5, 8, 8, 5}, 2)
	assert.Equal(t, (int)(ans), 6)
}

func Test_FindNode(t *testing.T) {
	//  1
	// / \
	// 0 2
	left := &solution.TreeNode{
		Val: 0,
	}
	right := &solution.TreeNode{
		Val: 2,
	}
	root := &solution.TreeNode{
		Val:   1,
		Left:  left,
		Right: right,
	}
	path := solution.FindNode(root, 2)
	assert.Equal(t, map[*solution.TreeNode]int{root: 0, right: 1}, path)
	path = solution.FindNode(root, 0)
	assert.Equal(t, map[*solution.TreeNode]int{root: 0, left: 1}, path)
	// add a new node
	secondLeaf := &solution.TreeNode{
		Val: 3,
	}
	right.Right = secondLeaf
	path = solution.FindNode(root, 3)
	assert.Equal(t, map[*solution.TreeNode]int{root: 0, right: 1, secondLeaf: 2}, path)
}

func Test_LowestCommonAncestor(t *testing.T) {
}
