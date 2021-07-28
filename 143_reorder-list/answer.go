/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reorder-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/21 12:16 上午
 */

package _reorder_list

import "github.com/s1m0n21/leetcode-go/utils"

func reorderList(head *utils.ListNode)  {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	left, right := head, slow.Next
	slow.Next = nil
	right = reverse(right)

	var leftTemp, rightTemp *utils.ListNode
	for left != nil && right != nil {
		leftTemp = left.Next
		rightTemp = right.Next

		left.Next = right
		left = leftTemp

		right.Next = left
		right = rightTemp
	}
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