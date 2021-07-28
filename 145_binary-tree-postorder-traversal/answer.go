/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/23 12:01 下午
 */

package _binary_tree_postorder_traversal

import "github.com/s1m0n21/leetcode-go/utils"

//func postorderTraversal(root *utils.TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	values := make([]int, 0)
//	values = append(values, postorderTraversal(root.Left)...)
//	values = append(values, postorderTraversal(root.Right)...)
//	values = append(values, root.Val)
//
//	return values
//}

func postorderTraversal(root *utils.TreeNode) []int {
	if root == nil {
		return nil
	}

	var stack = make([]*utils.TreeNode, 0)
	stack = append(stack, root.Left)
	stack = append(stack, root.Right)

	var res = make([]int, 0)
	res = append(res, root.Val)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)

		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
