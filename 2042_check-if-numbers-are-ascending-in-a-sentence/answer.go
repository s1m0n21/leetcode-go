/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/check-if-numbers-are-ascending-in-a-sentence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/20 11:05 上午
 */

package _check_if_numbers_are_ascending_in_a_sentence

import "strings"

func parseInt(s string) (int, bool) {
	n := 0
	for _, b := range s {
		if '0' <= b && b <= '9' {
			n = n * 10 + int(b-'0')
		} else {
			return 0, false
		}
	}
	return n, true
}

func areNumbersAscending(s string) bool {
	last := -1
	fields := strings.Fields(s)
	for _, s := range fields {
		if n, ok := parseInt(s); ok {
			if last != -1 && n <= last {
				return false
			} else {
				last = n
			}
		}
	}

	return true
}
