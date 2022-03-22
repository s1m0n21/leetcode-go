/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/22 11:18
 */

package _construct_binary_tree_from_preorder_and_inorder_traversal

import "github.com/s1m0n21/leetcode-go/utils"

func buildTree(preorder []int, inorder []int) *utils.TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &utils.TreeNode{Val: preorder[0]}
	var i int
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
