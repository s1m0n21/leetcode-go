/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/kth-smallest-element-in-a-bst/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/2 4:16 上午
 */

package _kth_smallest_element_in_a_bst

import "github.com/s1m0n21/leetcode-go/utils"

func kthSmallest(root *utils.TreeNode, k int) int {
	var stack []*utils.TreeNode
	var nums []int

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

	return nums[k-1]
}
