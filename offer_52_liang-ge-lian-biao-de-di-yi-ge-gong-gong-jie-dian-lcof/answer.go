/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/liang-ge-lian-biao-de-di-yi-ge-gong-gong-jie-dian-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/21 5:33 下午
 */

package offer_52_liang_ge_lian_biao_de_di_yi_ge_gong_gong_jie_dian_lcof

import "github.com/s1m0n21/leetcode-go/utils"

func getIntersectionNode(headA, headB *utils.ListNode) *utils.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}

		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}

	return p
}