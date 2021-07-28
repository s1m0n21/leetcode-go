/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sum-of-left-leaves/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/28 2:26 下午
 */

package _sum_of_left_leaves

import "github.com/s1m0n21/leetcode-go/utils"

func sumOfLeftLeaves(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*utils.TreeNode) int
	dfs = func(root *utils.TreeNode) int {
		res := 0
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				res += root.Left.Val
			} else {
				res += dfs(root.Left)
			}
		}

		if root.Right != nil {
			res += dfs(root.Right)
		}

		return res
	}

	return dfs(root)
}
