/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/add-two-numbers-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/19 12:07 上午
 */

package _add_two_numbers_ii

import "github.com/s1m0n21/leetcode-go/utils"

func addTwoNumbers(l1 *utils.ListNode, l2 *utils.ListNode) *utils.ListNode {
	var s1 = make([]int, 0)
	var s2 = make([]int, 0)

	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	carry := 0
	dummy := &utils.ListNode{}
	for len(s1) > 0 || len(s2) > 0 {
		var n1, n2 = 0, 0
		if len(s1) > 0 {
			n1 = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			n2 = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		sum := n1 + n2 + carry
		next := dummy.Next
		dummy.Next = &utils.ListNode{Val: sum % 10, Next: next}
		carry = sum / 10
	}

	if carry != 0 {
		next := dummy.Next
		dummy.Next = &utils.ListNode{Val: carry, Next: next}
	}

	return dummy.Next
}
