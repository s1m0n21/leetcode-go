/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/25 9:00 下午
 */

package _top_k_frequent_elements

import "testing"

func TestAnswer(t *testing.T) {
	nums := []int{1,1,1,2,2,3}
	k := 2

	t.Logf("answer = %+v", topKFrequent(nums, k))
}
