/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/path-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/27 3:12 下午
 */

package _path_sum

import "github.com/s1m0n21/leetcode-go/utils"

func hasPathSum(root *utils.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}

	return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}
