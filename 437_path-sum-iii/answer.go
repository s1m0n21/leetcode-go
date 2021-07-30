/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/path-sum-iii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/30 4:01 下午
 */

package _path_sum_iii

import "github.com/s1m0n21/leetcode-go/utils"

func pathSum(root *utils.TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := findPath(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)

	return res
}

func findPath(node *utils.TreeNode, target int) int {
	if node == nil {
		return 0
	}

	var res = 0
	if node.Val == target {
		res += 1
	}

	res += findPath(node.Left, target-node.Val)
	res += findPath(node.Right, target-node.Val)

	return res
}
