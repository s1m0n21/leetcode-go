/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/delete-node-in-a-bst/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/31 4:14 上午
 */

package _delete_node_in_a_bst

import (
	"github.com/s1m0n21/leetcode-go/utils"
)

func deleteNode(root *utils.TreeNode, key int) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}

	if root.Left == nil {
		return root.Right
	} else if root.Right == nil {
		return root.Left
	}

	min := root.Right
	for min.Left != nil {
		min = min.Left
	}
	root.Val = min.Val
	root.Right = delMinNode(root.Right)

	return root
}

func delMinNode(node *utils.TreeNode) *utils.TreeNode {
	if node.Left == nil {
		prev := node.Right
		node.Right = nil
		return prev
	}
	node.Left = delMinNode(node.Left)
	return node
}
