/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/symmetric-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/26 4:47 下午
 */

package _symmetric_tree

import "github.com/s1m0n21/leetcode-go/utils"

func isSymmetric(root *utils.TreeNode) bool {
	if root == nil {
		return true
	}

	var queue = make([]*utils.TreeNode, 0)
	queue = append(queue, root.Left, root.Right)

	for len(queue) > 0 {
		l, r := queue[0], queue[1]
		queue = queue[2:]

		if l == nil && r == nil {
			continue
		}

		if l == nil || r == nil || l.Val != r.Val {
			return false
		}

		queue = append(queue, l.Left, r.Right, l.Right, r.Left)
	}

	return true
}
