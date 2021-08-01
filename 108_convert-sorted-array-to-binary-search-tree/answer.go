/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/31 8:59 下午
 */

package _convert_sorted_array_to_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func sortedArrayToBST(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &utils.TreeNode{Val: nums[0]}
	}

	m := len(nums)/2
	root := &utils.TreeNode{Val: nums[m]}
	root.Left = sortedArrayToBST(nums[:m])
	root.Right = sortedArrayToBST(nums[m+1:])

	return root
}
