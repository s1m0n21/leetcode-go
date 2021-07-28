/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-linked-list-elements/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/5 7:35 下午
 */

package _remove_linked_list_elements

import "github.com/s1m0n21/leetcode-go/utils"

func removeElements(head *utils.ListNode, val int) *utils.ListNode {
	var out *utils.ListNode
	var tail *utils.ListNode

	for head != nil {
		if head.Val != val {
			if out == nil {
				out = &utils.ListNode{Val: head.Val}
				tail = out
			} else {
				tail.Next = &utils.ListNode{Val: head.Val}
				tail = tail.Next
			}
		}
		head = head.Next
	}

	return out
}
