/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/linked-list-cycle-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/6 11:09 下午
 */

package _linked_list_cycle_ii

import "github.com/s1m0n21/leetcode-go/utils"

func detectCycle(head *utils.ListNode) *utils.ListNode {
	var fast = head
	var slow = head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = head
	for slow != fast {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
