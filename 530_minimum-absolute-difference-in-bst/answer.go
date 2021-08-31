/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/minimum-absolute-difference-in-bst/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/31 4:20 下午
 */

package _minimum_absolute_difference_in_bst

import (
	"github.com/s1m0n21/leetcode-go/utils"
	"math"
)

func getMinimumDifference(root *utils.TreeNode) int {
	var travel func(node *utils.TreeNode)
	var prev *utils.TreeNode
	var min = math.MaxInt32

	travel = func(node *utils.TreeNode) {
		if node == nil {
			return
		}

		travel(node.Left)
		if prev != nil && node.Val - prev.Val < min {
			min = node.Val - prev.Val
		}
		prev = node
		travel(node.Right)
	}

	travel(root)

	return min
}
