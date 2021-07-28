/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/insertion-sort-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/19 10:33 下午
 */

package _insertion_sort_list

import "github.com/s1m0n21/leetcode-go/utils"

func insertionSortList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	var dummy = &utils.ListNode{Next: head}
	curr := head.Next
	last := head

	for curr != nil {
		if last.Val <= curr.Val {
			last = last.Next
		} else {
			prev := dummy
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			last.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = last.Next
	}

	return dummy.Next
}
