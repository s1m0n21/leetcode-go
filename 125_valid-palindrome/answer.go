/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/valid-palindrome/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/29 1:42 上午
 */

package _valid_palindrome

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	i, j := 0, len(s)-1
	for i < j {
		if !check(s[i]) {
			i++
			continue
		} else if !check(s[j]) {
			j--
			continue
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func check(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
