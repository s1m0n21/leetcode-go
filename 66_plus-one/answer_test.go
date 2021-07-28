/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/28 7:49 下午
 */

package _plus_one

import "testing"

func TestAnswer(t *testing.T) {
	digits := []int{8,9,9,9}

	t.Logf("answer = %+v", plusOne(digits))
}