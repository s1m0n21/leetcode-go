/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/linked-list-components/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/28 16:09
 */

package _linked_list_components

import "github.com/s1m0n21/leetcode-go/utils"

func numComponents(head *utils.ListNode, nums []int) int {
	var ret int

	r := make(map[int]struct{})
	for _, n := range nums {
		r[n] = struct{}{}
	}

	for head != nil {
		if _, in := r[head.Val]; in {
			ret++
			for head != nil {
				if _, in := r[head.Val]; !in {
					break
				}
				head = head.Next
			}
		}
		if head != nil {
			head = head.Next
		}
	}

	return ret
}
