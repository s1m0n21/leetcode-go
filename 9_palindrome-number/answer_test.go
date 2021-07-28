/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 11:21 上午
 */

package _palindrome_number

import "testing"

func TestAnswer(t *testing.T) {
	x := -121
	t.Logf("answer = %v", isPalindrome(x))
}
