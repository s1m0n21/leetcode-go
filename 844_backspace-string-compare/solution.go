/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/backspace-string-compare/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/25 10:51
 */

package _backspace_string_compare

func backspaceCompare(s string, t string) bool {
	var a, b []byte
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(a) > 0 {
				a = a[:len(a)-1]
			}
			continue
		}
		a = append(a, s[i])
	}

	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(b) > 0 {
				b = b[:len(b)-1]
			}
			continue
		}
		b = append(b, t[i])
	}

	return string(a) == string(b)
}
