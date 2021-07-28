/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/middle-of-the-linked-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/21 12:18 上午
 */

package _middle_of_the_linked_list

import "github.com/s1m0n21/leetcode-go/utils"

func middleNode(head *utils.ListNode) *utils.ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
