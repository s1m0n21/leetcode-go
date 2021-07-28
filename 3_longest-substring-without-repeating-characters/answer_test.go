/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/4 10:27 下午
 */

package _longest_substring_without_repeating_characters

import "testing"

func TestAnswer(t *testing.T) {
	s := "bcdeeefghjjck"
	t.Logf("answer = %d", lengthOfLongestSubstring(s))
}
