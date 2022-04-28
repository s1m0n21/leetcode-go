/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/shan-chu-lian-biao-de-jie-dian-lcof/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/27 20:04
 */

package offer_18_shan_chu_lian_biao_de_jie_dian_lcof

import "github.com/s1m0n21/leetcode-go/utils"

func deleteNode(head *utils.ListNode, val int) *utils.ListNode {
	dummy := &utils.ListNode{Next: head}
	curr := head
	prev := dummy

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	return dummy.Next
}
