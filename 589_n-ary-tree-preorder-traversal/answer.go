/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/23 14:01
 */

package _n_ary_tree_preorder_traversal

import "github.com/s1m0n21/leetcode-go/utils"

func preorder(root *utils.NTreeNode) []int {
	if root == nil {
		return nil
	}

	var stack []*utils.NTreeNode
	var nums []int

	stack = append(stack, root)

	for len(stack) > 0 {
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nums = append(nums, root.Val)

		for i := len(root.Children) - 1; i >= 0; i-- {
			stack = append(stack, root.Children[i])
		}
	}

	return nums
}
