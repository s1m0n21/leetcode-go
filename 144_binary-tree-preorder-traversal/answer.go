/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 11:42 上午
 */

package _binary_tree_preorder_traversal

import "github.com/s1m0n21/leetcode-go/utils"

//func preorderTraversal(root *utils.TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	values := make([]int, 0)
//	values = append(values, root.Val)
//	values = append(values, preorderTraversal(root.Left)...)
//	values = append(values, preorderTraversal(root.Right)...)
//
//	return values
//}

func preorderTraversal(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack = make([]*utils.TreeNode, 0)
	stack = append(stack, root.Right)
	stack = append(stack, root.Left)

	var res = make([]int, 0)
	res = append(res, root.Val)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)

		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}

	return res
}
