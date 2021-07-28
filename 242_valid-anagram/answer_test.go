/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/2 1:50 上午
 */

package _valid_anagram

import "testing"

func TestAnswer(t *testing.T) {
	s1 := "cat"
	s2 := "rat"

	t.Logf("answer = %t", isAnagram(s1, s2))
}
