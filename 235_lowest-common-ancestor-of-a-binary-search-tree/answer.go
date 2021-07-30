/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/30 5:08 下午
 */

package _lowest_common_ancestor_of_a_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}

	return root
}
