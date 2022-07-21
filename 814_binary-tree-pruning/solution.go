/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/binary-tree-pruning/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/21 09:39
 */

package _binary_tree_pruning

import "github.com/s1m0n21/leetcode-go/utils"

func pruneTree(root *utils.TreeNode) *utils.TreeNode {
	var dfs func(node *utils.TreeNode) bool

	dfs = func(node *utils.TreeNode) bool {
		if node.Left != nil && !dfs(node.Left) {
			node.Left = nil
		}
		if node.Right != nil && !dfs(node.Right) {
			node.Right = nil
		}
		if node.Left == nil && node.Right == nil {
			return node.Val == 1
		}
		return true
	}

	dfs(root)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		root = nil
	}

	return root
}
