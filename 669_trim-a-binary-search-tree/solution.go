/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/trim-a-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/9/10 12:49
 */

package _trim_a_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func trimBST(root *utils.TreeNode, low int, high int) *utils.TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)

	return root
}
