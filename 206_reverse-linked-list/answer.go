/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-linked-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/18 12:05 下午
 */

package _reverse_linked_list

import "github.com/s1m0n21/leetcode-go/utils"

func reverseList(head *utils.ListNode) *utils.ListNode {
	var prev *utils.ListNode
	var curr = head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
