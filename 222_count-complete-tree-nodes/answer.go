/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/count-complete-tree-nodes/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/26 6:03 下午
 */

package _count_complete_tree_nodes

import "github.com/s1m0n21/leetcode-go/utils"

func countNodes(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := 0, 0
	l, r := root, root

	for l != nil {
		lh++
		l = l.Left
	}

	for r != nil {
		rh++
		r = r.Right
	}

	if lh == rh {
		return 1 << lh - 1
	}

	return countNodes(root.Left) + countNodes(root.Right) + 1
}
