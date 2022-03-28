/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/convert-binary-number-in-a-linked-list-to-integer/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/28 16:23
 */

package _convert_binary_number_in_a_linked_list_to_integer

import "github.com/s1m0n21/leetcode-go/utils"

func getDecimalValue(head *utils.ListNode) int {
	var ans int

	for head != nil {
		ans = (ans << 1) | head.Val
		head = head.Next
	}

	return ans
}
