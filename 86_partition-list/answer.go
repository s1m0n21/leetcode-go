/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/partition-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/18 4:05 下午
 */

package _partition_list

import "github.com/s1m0n21/leetcode-go/utils"

func partition(head *utils.ListNode, x int) *utils.ListNode {
	if head == nil {
		return nil
	}

	dummyB := &utils.ListNode{}
	dummyS := &utils.ListNode{}
	bHead := dummyB
	sHead := dummyS

	for head != nil {
		if head.Val < x {
			dummyS.Next = head
			dummyS = head
		} else {
			dummyB.Next = head
			dummyB = head
		}
		head = head.Next
	}

	dummyS.Next = bHead.Next
	dummyB.Next = nil

	return sHead.Next
}
