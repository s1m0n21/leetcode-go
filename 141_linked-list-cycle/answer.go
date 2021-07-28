/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/linked-list-cycle/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/6 10:54 下午
 */

package _linked_list_cycle

import "github.com/s1m0n21/leetcode-go/utils"

func hasCycle(head *utils.ListNode) bool {
	var fast = head
	var slow = head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
