/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/lian-biao-zhong-dao-shu-di-kge-jie-dian-lcof/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/2 9:02 上午
 */

package offer_22_lian_biao_zhong_dao_shu_di_kge_jie_dian_lcof

import "github.com/s1m0n21/leetcode-go/utils"

func getKthFromEnd(head *utils.ListNode, k int) *utils.ListNode {
	if head == nil {
		return nil
	}

	var length int
	curr := head
	for curr != nil {
		curr = curr.Next
		length++
	}

	for i := 0; i < length-k; i++ {
		head = head.Next
	}

	return head
}
