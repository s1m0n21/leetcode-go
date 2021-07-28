/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/20 8:35 下午
 */

package _remove_nth_node_from_end_of_list

import "github.com/s1m0n21/leetcode-go/utils"

func removeNthFromEnd(head *utils.ListNode, n int) *utils.ListNode {
	var dummy = &utils.ListNode{Next: head}
	p, q := dummy, dummy
	for i := 0; i <= n; i++ {
		q = q.Next
	}

	for q != nil {
		p = p.Next
		q = q.Next
	}

	p.Next = p.Next.Next

	return dummy.Next
}
