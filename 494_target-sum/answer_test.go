/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 2:01 上午
 */

package _target_sum

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	target := 3
	t.Logf("answer = %d", findTargetSumWays(nums, target))
}
