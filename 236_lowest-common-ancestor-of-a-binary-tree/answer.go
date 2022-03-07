/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/7 11:14 AM
 */

package _lowest_common_ancestor_of_a_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
