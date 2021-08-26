/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/23 9:41 上午
 */

package _intersection_of_two_linked_lists

import "github.com/s1m0n21/leetcode-go/utils"

func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	if headA == nil && headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}

	return p
}
