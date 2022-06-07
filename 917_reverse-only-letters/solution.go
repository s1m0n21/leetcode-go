/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/reverse-only-letters/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/7 09:31
 */

package _reverse_only_letters

func reverseOnlyLetters(s string) string {
	bytes := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isLetter(bytes[l]) {
			l++
		}

		for l < r && !isLetter(bytes[r]) {
			r--
		}

		bytes[l], bytes[r] = bytes[r], bytes[l]
		r--
		l++
	}

	return string(bytes)
}

func isLetter(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}
