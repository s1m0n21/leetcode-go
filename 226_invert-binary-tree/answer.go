/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/invert-binary-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/26 3:47 下午
 */

package _invert_binary_tree

import "github.com/s1m0n21/leetcode-go/utils"

func invertTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
