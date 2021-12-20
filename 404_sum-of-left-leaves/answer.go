/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sum-of-left-leaves/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/28 2:26 下午
 */

package _sum_of_left_leaves

import "github.com/s1m0n21/leetcode-go/utils"

func sumOfLeftLeaves(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	var sum int
	var stack []*utils.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
				sum += root.Left.Val
			}
			root = root.Left
		}
		root = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}

	return sum
}
