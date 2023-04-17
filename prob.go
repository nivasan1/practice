package prob

import (
	"fmt"
	_ "fmt"
)

// maximal values
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// given integers a_i, return b_i where b_i = \prod_{j \in I \ {i}} b_j,
// cannot use division for problems, must be O(n) time
func ProductExceptSelf(nums []int) []int {
	// initialize array vals, set to 0
	lessIdx := make([]int, len(nums))
	greaterIdx := make([]int, len(nums))
	// iterate over nums
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			lessIdx[i] = 1
			greaterIdx[len(nums)-(i+1)] = 1
			continue
		}
		lessIdx[i] = lessIdx[i-1] * nums[i-1]
		greaterIdx[len(nums)-(i+1)] = greaterIdx[len(nums)-(i)] * nums[len(nums)-i]

	}

	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = lessIdx[i] * greaterIdx[i]
	}
	return ans
}

// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
func MaxSubArray(nums []int) int {
	sum := nums[0]
	FinalSum := sum
	// iterate through the array,
	// reset sum to 0, when sum is less than 0
	// reset start idx as well
	for i := 1; i < len(nums); i++ {
		if sum > FinalSum {
			FinalSum = sum
		}
		// if the sum is less than 0, restart
		if sum <= 0 {
			sum = 0
		}
		// add next item to sum
		sum += nums[i]
	}
	return FinalSum
}

// can also find the maximal sum-recursively, by splitting array into two 
// ^^ this is in-correct as it ignores the cases where the left max-sum is dimimished by larger negative in-between
func MaxSubArray2(nums []int) int {
	low , high, sum := maxSubArray(nums, 0, len(nums) - 1)
	fmt.Println(low, high)
	return sum
}

func maxSubArray(nums []int, low, high int) (int, int, int) {
	// base-case, if they are equal return 
	if low == high {
		return low, high, nums[low]
	}
	mid := (low + high) / 2
	leftLowBound, leftUpBound, sumLeft := maxSubArray(nums, low, mid)
	rightLowBound, rightUpBound, sumRight := maxSubArray(nums, mid + 1, high)
	// determine if there is a link between the two arrays, i.e the sum of
	link := findLink(nums,leftUpBound, rightLowBound)
	total := link + sumLeft + sumRight
	if total >= sumLeft && total >= sumRight {
		return leftLowBound, rightUpBound, total
	}

	if sumLeft > sumRight {
		return leftLowBound, leftUpBound, sumLeft
	}
	return rightLowBound, rightUpBound, sumRight
}

func findLink(nums []int, leftBound, rightBound int) int {
	sum := 0
	for i := leftBound + 1; i < rightBound; i++ {
		sum += nums[i]
	}
	return sum
}

// find the minimum of a sorted array that is rotated
// use binary search to
func FindMin(nums []int) int {
	start := 0
	end := len(nums) - 1
	mid := (start + end) / 2
	if nums[start] <= nums[end] {
		return nums[start]
	}
	for {
		// we have two elements, one is the pivot
		if end == start {
			if nums[end] < nums[(end+1)%len(nums)] {
				return nums[end]
			}
			return nums[end+1]
		}
		// first half is in sorted order, elements are unique, never equal
		if nums[start] < nums[mid] {
			start = mid
		} else {
			// the first half is not in sorted order
			end = mid
		}
		mid = (start + end) / 2
	}
}

func Search(nums []int, target int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	// distance to cover in the array
	dist := ((len(nums) - 1) / 2) + 1
	i := dist
	steps := 0
	for {
		//  we have found the number
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			if i >= dist {
				i = i - dist
			} else {
				i = len(nums) - (dist - i)
			}
		} else if nums[i] < target {
			i = (i + dist) % len(nums)
		}
		steps++
		if steps > (len(nums)/2)+1 {
			return -1
		}
		// divide dist by 2
		if dist != 1 {
			dist = dist >> 1
		}
	}
}

func ClimbStairs(n int) int {
	val := make([]int, n+1)
	val[0] = 0
	val[1] = 1
	val[2] = 2
	for i := 3; i < len(val); i++ {
		val[i] = val[i-1] + val[i-2]
	}
	return val[len(val)-1]
}

// solve problem recursively
func CoinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	amounts := make(map[int]int)
	for _, coin := range coins {
		// skip everything and return 1
		if coin == amount {
			return 1
		}
		// skip this coin
		if amount < coin {
			amounts[coin] = -1
			continue
		}
		change := CoinChange(coins, amount-coin)
		if change == -1 {
			amounts[coin] = -1
			continue
		}
		amounts[coin] = change + 1
	}

	min := MaxInt
	// if -1 is val for any of the amounts, it will be returned, as no value could be less
	for _, val := range amounts {
		// ignore
		if val == -1 {
			continue
		}
		if val < min {
			min = val
		}
	}
	if min == MaxInt {
		return -1
	}
	return min
}

// Solve dynamically
func CoinChange2(coins []int, amount int) int {
	amounts := make([]int, amount+1)
	vals := make([]int, len(coins))
	for amount := range amounts {
		if amount == 0 {
			amounts[0] = 0
			continue
		}
		for idx, coin := range coins {
			if coin > amount {
				vals[idx] = -1
				continue
			}
			vals[idx] = amounts[amount-coin]
		}

		min := MaxInt
		// if -1 is val for any of the amounts, it will be returned, as no value could be less
		for _, val := range vals {
			// ignore
			if val == -1 {
				continue
			}
			if val < min {
				min = val
			}
		}
		if min == MaxInt {
			amounts[amount] = -1
			continue
		}
		amounts[amount] = min + 1
	}
	return amounts[len(amounts)-1]
}

func LengthOfLIS(nums []int) int {
	vals := make(map[int]int)
	for _, num := range nums {
		if _, ok := vals[num]; !ok {
			vals[num] = 1
		}
		high := 0
		for highest, len := range vals {
			if num > highest {
				if len > high {
					high = len
				}
			}
			// pruning the search space
			if num == highest+1 {
				delete(vals, highest)
			}
		}
		vals[num] = high + 1
	}

	max := 0
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func LongestCommonSubsequence(text1 string, text2 string) int {
	small := ""
	large := ""
	if len(text1) < len(text2) {
		small = text1
		large = text2
	} else {
		small = text2
		large = text1
	}
	// map of individual characters to indices
	wordIdx := make(map[byte]int)
	for i := 0; i < len(small); i++ {
		for j := 0; j < len(large); j++ {
			// only record index of earliest occurance
			if _, ok := wordIdx[small[i]]; ok {
				continue
			}
			// the characters are equal, record idx
			if []byte(large)[j] == []byte(small)[i] {
				wordIdx[byte(small[i])] = j
				continue
			}
		}
	}
	subs_len := 0
	prevIdx := 0
	// iterate through characters of small
	for i := 0; i < len(small); i++ {
		// this letter is not present in larger word
		if _, ok := wordIdx[small[i]]; !ok {
			continue
		}
		if subs_len == 0 {
			subs_len = 1
			prevIdx = wordIdx[small[i]]
			continue
		}
		if wordIdx[small[i]] > prevIdx {
			subs_len++
		} else {
			subs_len = 0
		}
		prevIdx = wordIdx[small[i]]
	}

	return subs_len
}

func TaskSchedulerII(tasks []int, space int) int64 {
	active_tasks := make(map[int]int) /* is task currently being executed, (task type, index last encountered) */
	days := 0
	for _, task := range tasks {
		days++
		fmt.Printf("active_tasks: %v, days: %d\n", active_tasks, days)
		// this type task is being executed
		if _, ok := active_tasks[task]; ok {
			// replace the task with a new one
			delay := (space - (days - active_tasks[task])) + 1
			days += delay
			active_tasks[task] = days
		} else {
			// this type of task is not being executed
			active_tasks[task] = days
		}
	}

	return (int64)(days)
}

func minPathSum(grid [][]int) int {
	sumGrid := make([][]int, 0)
	// initialize sub-grid
	for i, subgrid := range grid {
		sumGrid[i] = make([]int, len(subgrid))
	}

	return 0
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// find p in the BST
	pathForP := FindNode(root, p.Val)
	// find q in the BST
	pathForQ := FindNode(root, q.Val)
	// find the smallest value from a node contained in both paths
	lowestHeight := MinInt
	var lca *TreeNode
	for node, heightOfNode := range pathForQ {
		// node exists in both paths
		if _, ok := pathForP[node]; ok {
			if heightOfNode > lowestHeight {
				lowestHeight = heightOfNode
				lca = node
			}
		}
	}
	return lca
}

// function used to find a node in the BST with value val
// return the list of nodes encountered along the way
func FindNode(root *TreeNode, val int) (path map[*TreeNode]int) {
	path = make(map[*TreeNode]int)
	depth := 0
	// keep going down tree, as long as there are valid children
	for root != nil {
		path[root] = depth
		depth++
		// if the desired value is less than the value of the current node
		// search to the left
		if val < root.Val {
			// add node to the path
			root = root.Left
			continue
		}
		// if it is greater search to the right
		if val > root.Val {
			root = root.Right
			continue
		}
		// they are equal return
		return
	}
	// node does not exist in tree exit
	return nil
}
