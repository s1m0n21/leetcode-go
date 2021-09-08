/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/split-a-string-in-balanced-strings/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/7 1:33 下午
 */

package _split_a_string_in_balanced_strings

func balancedStringSplit(s string) int {
	var res int
	l, r := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			l++
		} else if s[i] =='R' {
			r++
		}

		if l != 0 && r != 0 && l == r {
			res++
			l, r = 0, 0
		}
	}

	return res
}
