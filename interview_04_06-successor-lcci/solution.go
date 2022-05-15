/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/successor-lcci/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/16 02:32
 */

package interview_04_06_successor_lcci

import "github.com/s1m0n21/leetcode-go/utils"

func inorderSuccessor(root *utils.TreeNode, p *utils.TreeNode) *utils.TreeNode {
	var stack []*utils.TreeNode
	var prev *utils.TreeNode

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev == p {
			return root
		}
		prev = root
		root = root.Right
	}

	return nil
}
