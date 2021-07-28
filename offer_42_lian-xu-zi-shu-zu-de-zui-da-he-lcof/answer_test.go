/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/17 2:42 上午
 */

package offer_42_lian_xu_zi_shu_zu_de_zui_da_he_lcof

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}

	t.Logf("answer = %d", maxSubArray(nums))
}
