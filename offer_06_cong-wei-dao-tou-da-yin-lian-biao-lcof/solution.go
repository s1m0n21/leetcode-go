/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/27 19:55
 */

package offer_06_cong_wei_dao_tou_da_yin_lian_biao_lcof

import "github.com/s1m0n21/leetcode-go/utils"

func reversePrint(head *utils.ListNode) []int {
	var ret []int

	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	n := len(ret)
	for i := 0; i < n/2; i++ {
		ret[i], ret[n-i-1] = ret[n-i-1], ret[i]
	}

	return ret
}
