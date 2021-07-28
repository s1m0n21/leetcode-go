/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sort-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/20 11:49 上午
 */

package _sort_list

import "github.com/s1m0n21/leetcode-go/utils"

func sortList(head *utils.ListNode) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, fast, slow := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil

	l := sortList(head)
	r := sortList(slow)

	return merge(l, r)
}

func merge(l1, l2 *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	prev := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}

	return dummy.Next
}