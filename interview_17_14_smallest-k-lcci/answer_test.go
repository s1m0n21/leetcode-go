/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/26 11:35 上午
 */

package interview_17_14_smallest_k_lcci

import "testing"

func TestAnswer(t *testing.T) {
	arr := []int{1,3,5,7,2,4,6,8}
	k := 4

	t.Logf("answer = %+v", smallestK(arr, k))
}
