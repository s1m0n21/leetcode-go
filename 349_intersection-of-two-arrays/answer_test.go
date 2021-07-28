/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/30 11:35 下午
 */

package _intersection_of_two_arrays

import "testing"

func TestAnswer(t *testing.T) {
	n1 := []int{4,9,5}
	n2 := []int{9,4,9,8,4}

	t.Logf("answer = %+v", intersection(n1, n2))
}
