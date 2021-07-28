/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sum-root-to-leaf-numbers/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/27 4:02 下午
 */

package _sum_root_to_leaf_numbers

import "github.com/s1m0n21/leetcode-go/utils"

func sumNumbers(root *utils.TreeNode) int {
	var dfs func(*utils.TreeNode, int) int
	dfs = func(root *utils.TreeNode, prev int) int {
		if root == nil {
			return 0
		}

		sum := prev * 10 + root.Val
		if root.Left == nil && root.Right == nil {
			return sum
		}

		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}

	return dfs(root, 0)
}
