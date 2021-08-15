/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/valid-parenthesis-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/15 11:44 下午
 */

package _valid_parenthesis_string

func checkValidString(s string) bool {
	if len(s) == 0 {
		return true
	}

	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			r++
		} else {
			l++
		}

		if r > l {
			return false
		}
	}

	l, r = 0, 0
	for i := len(s)-1; i >= 0; i-- {
		if s[i] == '(' {
			l++
		} else {
			r++
		}

		if l > r {
			return false
		}
	}

	return true
}
