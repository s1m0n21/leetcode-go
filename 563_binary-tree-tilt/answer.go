/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/binary-tree-tilt/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/18 1:45 下午
 */

package _binary_tree_tilt

import "github.com/s1m0n21/leetcode-go/utils"

func findTilt(root *utils.TreeNode) int {
	var ans int
	var dfs func(root *utils.TreeNode) int

	dfs = func(root *utils.TreeNode) int {
		if root == nil {
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)
		ans += abs(left - right)

		return left + right + root.Val
	}
	dfs(root)

	return ans
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -1 * n
}
