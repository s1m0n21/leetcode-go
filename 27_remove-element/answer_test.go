/**
 * @Author         : s1m0n21
 * @Description    : Test Answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/25 2:58 下午
 */

package _remove_element

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3

	t.Logf("answer = %d, nums = %+v", removeElement(nums, val), nums)
}
