/**
 * @Author         : s1m0n21
 * @Description    : Answer test
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 2:07 下午
 */

package _minimum_size_subarray_sum

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{2,3,1,2,4,3}
	target := 7
	t.Logf("answer = %d", minSubArrayLen(target, nums))
}
