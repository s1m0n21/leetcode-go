/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/26 11:36 下午
 */

package _merge_sorted_array

import "testing"

func TestAnswer(t *testing.T) {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}

	merge(nums1, 3, nums2, 3)

	t.Logf("answer = %+v", nums1)
}
