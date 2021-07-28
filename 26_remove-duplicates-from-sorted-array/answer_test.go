/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/25 4:07 下午
 */

package _remove_duplicates_from_sorted_array

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{}

	t.Logf("answer = %d, nums = %+v", removeDuplicates(nums), nums)
}
