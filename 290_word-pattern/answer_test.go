/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/7/2 2:55 上午
 */

package _word_pattern

import "testing"

func TestAnswer(t *testing.T) {
	pattern := "abba"
	str := "dog cat cat dog"

	t.Logf("answer = %t", wordPattern(pattern, str))
}
