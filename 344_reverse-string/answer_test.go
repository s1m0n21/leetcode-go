/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 2:14 上午
 */

package _reverse_string

import "testing"

func TestAnswer(t *testing.T) {
	s := []byte("hello")
	reverseString(s)
	t.Logf("answer = %s", s)
}
