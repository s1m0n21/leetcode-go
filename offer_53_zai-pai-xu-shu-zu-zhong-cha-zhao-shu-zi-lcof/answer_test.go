/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/16 9:39 上午
 */

package offer_53_zai_pai_xu_shu_zu_zhong_cha_zhao_shu_zi_lcof

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{5,7,7,8,8,10}
	target := 5

	t.Logf("answer = %d", search(nums, target))
}
