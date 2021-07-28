/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/same-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/26 4:22 下午
 */

package _same_tree

import "github.com/s1m0n21/leetcode-go/utils"

func isSameTree(p *utils.TreeNode, q *utils.TreeNode) bool {
	if p == nil && q != nil || q == nil && p != nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	if p.Val != q.Val || !isSameTree(p.Left, q.Left) || !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}
