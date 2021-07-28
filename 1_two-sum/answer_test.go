/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 1:32 上午
 */

package _two_sum

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{3,2,4}
	target := 6

	t.Logf("answer = %v", twoSum(nums, target))
}
