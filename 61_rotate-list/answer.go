/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/rotate-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/20 9:31 下午
 */

package _rotate_list

import "github.com/s1m0n21/leetcode-go/utils"

func rotateRight(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := 1
	iter := head
	for iter.Next != nil {
		iter = iter.Next
		n++
	}

	diff := n - k % n
	if diff == n {
		return head
	}

	iter.Next = head
	for diff > 0 {
		iter = iter.Next
		diff--
	}

	res := iter.Next
	iter.Next = nil

	return res
}
