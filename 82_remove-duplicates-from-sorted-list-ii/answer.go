/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/19 1:29 下午
 */

package _remove_duplicates_from_sorted_list_ii

import "github.com/s1m0n21/leetcode-go/utils"

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	var dummy = &utils.ListNode{Val: -101}
	dummy.Next = head

	pre := dummy
	curr := dummy
	for curr != nil {
		temp := curr
		for temp.Next != nil && temp.Next.Val == temp.Val {
			temp = temp.Next
		}

		curr = temp.Next
		if curr == nil {
			pre.Next = nil
			break
		}

		if curr.Next != nil && curr.Next.Val == curr.Val {
			continue
		}

		pre.Next = curr
		pre = curr
	}

	return dummy.Next
}
