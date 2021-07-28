/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/6 1:44 下午
 */

package _four_sum_ii

import "testing"

func TestAnswer(t *testing.T) {
	nums1 := []int{1,2}
	nums2 := []int{-2,-1}
	nums3 := []int{-1,2}
	nums4 := []int{0,2}

	t.Logf("answer = %d", fourSumCount(nums1, nums2, nums3, nums4))
}
