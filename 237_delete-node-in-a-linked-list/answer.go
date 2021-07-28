/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/delete-node-in-a-linked-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/20 3:48 下午
 */

package _delete_node_in_a_linked_list

import "github.com/s1m0n21/leetcode-go/utils"

func deleteNode(node *utils.ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
