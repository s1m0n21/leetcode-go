/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/23 11:49 AM
 */

package _two_sum_iv_input_is_a_bst

import "github.com/s1m0n21/leetcode-go/utils"

func findTarget(root *utils.TreeNode, k int) bool {
	var nums []int
	var stack []*utils.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)
		root = root.Right
	}

	var l, r = 0, len(nums) - 1
	for l < r {
		sum := nums[l] + nums[r]
		if sum > k {
			r--
		} else if sum < k {
			l++
		} else {
			return true
		}
	}

	return false
}
