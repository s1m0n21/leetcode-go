/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/29 2:00 上午
 */

package _valid_palindrome

import "testing"

func TestAnswer(t *testing.T) {
	s := "race a car"
	t.Logf("answer = %v", isPalindrome(s))
}
