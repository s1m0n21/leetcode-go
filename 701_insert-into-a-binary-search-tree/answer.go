/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/23 12:20 PM
 */

package _insert_into_a_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func insertIntoBST(root *utils.TreeNode, val int) *utils.TreeNode {
	if root == nil {
		return &utils.TreeNode{Val: val}
	}

	node := root
	for node != nil {
		if val < node.Val {
			if node.Left == nil {
				node.Left = &utils.TreeNode{Val: val}
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = &utils.TreeNode{Val: val}
				break
			}
			node = node.Right
		}
	}

	return root
}
