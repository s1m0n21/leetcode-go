/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 11:54 上午
 */

package _binary_tree_inorder_traversal

import "github.com/s1m0n21/leetcode-go/utils"

//func inorderTraversal(root *utils.TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	values := make([]int, 0)
//	values = append(values, inorderTraversal(root.Left)...)
//	values = append(values, root.Val)
//	values = append(values, inorderTraversal(root.Right)...)
//
//	return values
//}

func inorderTraversal(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack = make([]*utils.TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	var res = make([]int, 0)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}

	return res
}
