/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/find-bottom-left-tree-value/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/22 10:09
 */

package _find_bottom_left_tree_value

import "github.com/s1m0n21/leetcode-go/utils"

func findBottomLeftValue(root *utils.TreeNode) int {
	var ret int
	queue := []*utils.TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		ret = node.Val
	}
	return ret
}
