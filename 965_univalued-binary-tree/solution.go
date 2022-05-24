/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/univalued-binary-tree/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/24 09:19
 */

package _univalued_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func isUnivalTree(root *utils.TreeNode) bool {
	n := root.Val

	var queue []*utils.TreeNode
	queue = append(queue, root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Val != n {
			return false
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return true
}
