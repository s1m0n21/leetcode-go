/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 2:23 上午
 */

package _reverse_vowels_of_a_string

import "testing"

func TestAnswer(t *testing.T) {
	s := "hello"
	t.Logf("answer = %s", reverseVowels(s))
}
