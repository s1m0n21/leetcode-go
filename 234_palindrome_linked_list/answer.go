/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/palindrome-linked-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/21 12:46 上午
 */

package _palindrome_linked_list

import "github.com/s1m0n21/leetcode-go/utils"

func isPalindrome(head *utils.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left, right := head, slow.Next
	slow.Next = nil

	right = reverse(right)

	for left != nil && right != nil {
		if left.Val != right.Val {
			return false
		}

		left = left.Next
		right = right.Next
	}

	return true
}

func reverse(head *utils.ListNode) *utils.ListNode {
	var pre *utils.ListNode = nil
	var curr = head

	for curr != nil {
		next := curr.Next

		curr.Next = pre
		pre = curr
		curr = next
	}

	return pre
}
