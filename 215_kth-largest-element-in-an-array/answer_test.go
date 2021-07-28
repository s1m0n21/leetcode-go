/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/28 11:29 下午
 */

package _kth_largest_element_in_an_array

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{3,2,3,1,2,4,5,5,6}
	k := 4

	t.Logf("answer = %d", findKthLargest(nums, k))
}
