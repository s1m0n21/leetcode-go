/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/30 1:50 上午
 */

package _find_all_anagrams_in_a_string

import "testing"

func TestAnswer(t *testing.T) {
	s := "abacbabc"
	p := "abc"

	t.Logf("answer = %+v", findAnagrams(s, p))
}
