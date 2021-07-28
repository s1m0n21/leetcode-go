/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 12:07 上午
 */

package _two_sum_ii_input_array_is_sorted

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{-1, 0}
	target := -1
	t.Logf("answer = %+v", twoSum(nums, target))
}
