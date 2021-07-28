/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-nodes-in-k-group/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/19 5:33 下午
 */

package _reverse_nodes_in_k_group

import "github.com/s1m0n21/leetcode-go/utils"

func reverseKGroup(head *utils.ListNode, k int) *utils.ListNode {
	if k <= 1 {
		return head
	}

	var dummy = &utils.ListNode{Next: head}
	pre := dummy

loop:
	for {
		var next = pre
		for i := 0; i < k; i++ {
			if next.Next == nil {
				break loop
			}
			next = next.Next
		}

		node1, node2 := pre.Next, pre.Next.Next
		curr := node1
		pre.Next = next
		node1.Next = next.Next
		for i := 2; i <= k; i++ {
			temp := node2.Next
			node2.Next = node1
			node1, node2 = node2, temp
		}

		pre = curr
		next = curr
	}

	return dummy.Next
}
