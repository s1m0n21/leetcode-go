/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/30 10:02
 */

package _sum_of_root_to_leaf_binary_numbers

import "github.com/s1m0n21/leetcode-go/utils"

func sumRootToLeaf(root *utils.TreeNode) int {
	var sum int
	var dfs func(node *utils.TreeNode, i int)

	dfs = func(node *utils.TreeNode, i int) {
		i |= node.Val

		if node.Left == nil && node.Right == nil {
			sum += i
			return
		}

		if node.Left != nil {
			dfs(node.Left, i << 1)
		}

		if node.Right != nil {
			dfs(node.Right, i << 1)
		}
	}

	dfs(root, 0)

	return sum
}
