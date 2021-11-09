/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binode-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/9 5:07 下午
 */

package interview_17_12_binode_lcci

import "github.com/s1m0n21/leetcode-go/utils"

func convertBiNode(root *utils.TreeNode) *utils.TreeNode {
	var stack []*utils.TreeNode
	var dummy = &utils.TreeNode{}
	var tail = dummy

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root.Left = nil
		tail.Right = root
		root = root.Right
		tail = tail.Right
	}

	return dummy.Right
}
