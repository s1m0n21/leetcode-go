/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/validate-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/30 5:29 下午
 */

package _validate_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func isValidBST(root *utils.TreeNode) bool {
	return checkBST(root, 1<<63 - 1, -1 << 63)
}

func checkBST(root *utils.TreeNode, upper, lower int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return checkBST(root.Left, root.Val, lower) && checkBST(root.Right, upper, root.Val)
}
