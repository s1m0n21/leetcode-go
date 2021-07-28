/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/odd-even-linked-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/18 4:41 下午
 */

package _odd_even_linked_list

import "github.com/s1m0n21/leetcode-go/utils"

func oddEvenList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	dummyO := &utils.ListNode{}
	dummyE := &utils.ListNode{}
	oddHead := dummyO
	evenHead := dummyE

	i := 0
	for head != nil {
		if i % 2 == 0 {
			dummyE.Next = head
			dummyE = head
		} else {
			dummyO.Next = head
			dummyO = head
		}
		head = head.Next
		i++
	}

	dummyE.Next = oddHead.Next
	dummyO.Next = nil

	return evenHead.Next
}
