/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/18 3:35 下午
 */

package _remove_duplicates_from_sorted_list

import "github.com/s1m0n21/leetcode-go/utils"

func deleteDuplicates(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}

	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}
