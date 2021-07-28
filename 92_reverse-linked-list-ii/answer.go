/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-linked-list-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/18 1:32 下午
 */

package _reverse_linked_list_ii

import "github.com/s1m0n21/leetcode-go/utils"

func reverseBetween(head *utils.ListNode, left int, right int) *utils.ListNode {
	dummy := &utils.ListNode{Val: -1}
	dummy.Next = head
	pre := dummy
	for i := 0; i < left - 1; i++ {
		pre = pre.Next
	}

	curr := pre.Next
	for i := 0; i < right - left; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}

	return dummy.Next
}
