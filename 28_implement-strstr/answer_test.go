/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/8 11:44 上午
 */

package _implement_strstr

import "testing"

func TestAnswer(t *testing.T) {
	haystack := "aaaabaaac"
	needle := "aaac"

	t.Logf("answer = %d, haystack = %s, needle = %s", strStr(haystack, needle), haystack, needle)
}
