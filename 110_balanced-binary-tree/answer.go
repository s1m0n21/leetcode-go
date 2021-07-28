/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/balanced-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/27 4:21 下午
 */

package _balanced_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func isBalanced(root *utils.TreeNode) bool {
	return height(root) >= 0
}

func height(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	lh := height(root.Left)
	rh := height(root.Right)
	if lh == -1 || rh == -1 || abs(lh - rh) > 1 {
		return -1
	}

	return max(lh, rh) + 1
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
