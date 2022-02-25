/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/path-sum-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/29 10:54 下午
 */

package _path_sum_ii

import "github.com/s1m0n21/leetcode-go/utils"

func pathSum(root *utils.TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}

	var res = make([][]int, 0)
	var path = make([]int, 0)
	var dfs func(*utils.TreeNode, int)

	dfs = func(node *utils.TreeNode, target int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()

		if node.Left == nil && node.Right == nil && target-node.Val == 0 {
			res = append(res, append([]int(nil), path...))
		}

		dfs(node.Left, target-node.Val)
		dfs(node.Right, target-node.Val)
	}
	dfs(root, targetSum)

	return res
}
