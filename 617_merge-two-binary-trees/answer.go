/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/merge-two-binary-trees/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/24 09:33
 */

package _merge_two_binary_trees

import "github.com/s1m0n21/leetcode-go/utils"

func mergeTrees(root1 *utils.TreeNode, root2 *utils.TreeNode) *utils.TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}

	node := &utils.TreeNode{}
	if root1 != nil && root2 != nil {
		node.Val = root1.Val + root2.Val
		node.Left = mergeTrees(root1.Left, root2.Left)
		node.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 != nil {
		return root1
	} else if root2 != nil {
		return root2
	}

	return node
}
