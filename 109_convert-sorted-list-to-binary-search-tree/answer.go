/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/4 1:18 上午
 */

package _convert_sorted_list_to_binary_search_tree

import "github.com/s1m0n21/leetcode-go/utils"

func sortedListToBST(head *utils.ListNode) *utils.TreeNode {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	return makeBST(nums)
}

func makeBST(nums []int) *utils.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &utils.TreeNode{Val: nums[0]}
	}

	m := len(nums)/2
	root := &utils.TreeNode{Val: nums[m]}
	root.Left = makeBST(nums[:m])
	root.Right = makeBST(nums[m+1:])

	return root
}
