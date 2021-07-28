/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/merge-two-sorted-lists/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/19 4:13 下午
 */

package _merge_two_sorted_lists

import "github.com/s1m0n21/leetcode-go/utils"

func mergeTwoLists(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var dummy = &utils.ListNode{}

	curr := dummy
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				curr.Next = l1
				l1 = l1.Next
			} else {
				curr.Next = l2
				l2 = l2.Next
			}
			curr = curr.Next
			continue
		}

		if l1 == nil {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}

		curr = curr.Next
	}

	curr.Next = nil

	return dummy.Next
}
