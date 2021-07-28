/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/26 1:35 上午
 */

package _remove_duplicates_from_sorted_array_ii

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3, 3, 3}
	t.Logf("answer = %d, nums = %+v", removeDuplicates(nums), nums)
}
