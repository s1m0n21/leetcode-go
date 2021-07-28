/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/merge-k-sorted-lists/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/25 11:41 下午
 */

package _merge_k_sorted_lists

import "github.com/s1m0n21/leetcode-go/utils"

func mergeKLists(lists []*utils.ListNode) *utils.ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	return mergeList(mergeKLists(lists[len(lists)/2:]), mergeKLists(lists[:len(lists)/2]))
}

func mergeList(a, b *utils.ListNode) *utils.ListNode {
	var dummy = &utils.ListNode{}
	curr := dummy
	for a != nil && b != nil {
		if a.Val < b.Val {
			curr.Next = a
			a = a.Next
		} else {
			curr.Next = b
			b = b.Next
		}
		curr = curr.Next
	}

	for a != nil {
		curr.Next = a
		a = a.Next
		curr = curr.Next
	}

	for b != nil {
		curr.Next = b
		b = b.Next
		curr = curr.Next
	}

	curr.Next = nil

	return dummy.Next
}
