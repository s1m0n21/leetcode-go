/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/merge-nodes-in-between-zeros/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/5 10:01
 */

package _merge_nodes_in_between_zeros

import "github.com/s1m0n21/leetcode-go/utils"

func mergeNodes(head *utils.ListNode) *utils.ListNode {
	dummy := &utils.ListNode{}
	tail := dummy

	n := 0
	for head != nil {
		if n > 0 && head.Val == 0 {
			tail.Next = &utils.ListNode{Val: n}
			tail = tail.Next
			n = 0
		} else {
			n += head.Val
		}
		head = head.Next
	}

	return dummy.Next
}
