/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/remove-duplicate-node-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/7 12:03 PM
 */

package interview_02_01_remove_duplicate_node_lcci

import "github.com/s1m0n21/leetcode-go/utils"

func removeDuplicateNodes(head *utils.ListNode) *utils.ListNode {
	var dummy = &utils.ListNode{Next: head}
	var record = make(map[int]int)

	prev := dummy
	curr := dummy.Next
	for curr != nil {
		if record[curr.Val] > 0 {
			prev.Next = curr.Next
		} else {
			record[curr.Val]++
			prev = curr
		}
		curr = curr.Next
	}

	return dummy.Next
}
