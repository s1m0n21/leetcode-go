/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/second-minimum-node-in-a-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/27 11:07 上午
 */

package _second_minimum_node_in_a_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func findSecondMinimumValue(root *utils.TreeNode) int {
	res := -1
	rv := root.Val

	var dfs func(*utils.TreeNode)
	dfs = func(node *utils.TreeNode) {
		if node == nil || res != -1 && node.Val >= res {
			return
		}
		if node.Val > rv {
			res = node.Val
		}
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return res
}
