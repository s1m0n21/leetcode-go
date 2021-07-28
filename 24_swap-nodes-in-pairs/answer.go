/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/swap-nodes-in-pairs/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/19 4:30 下午
 */

package _swap_nodes_in_pairs

import "github.com/s1m0n21/leetcode-go/utils"

func swapPairs(head *utils.ListNode) *utils.ListNode {
	var dummy = &utils.ListNode{}
	dummy.Next = head
	pre := dummy

	for pre.Next != nil && pre.Next.Next != nil {
		node1, node2 := pre.Next, pre.Next.Next
		node1.Next = node2.Next
		node2.Next = node1
		pre.Next = node2
		pre = node1
	}

	return dummy.Next
}
