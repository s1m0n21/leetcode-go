/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/26 9:40 上午
 */

package _search_in_a_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func searchBST(root *utils.TreeNode, val int) *utils.TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else if root.Val < val  {
			root = root.Right
		}
	}

	return nil
}
