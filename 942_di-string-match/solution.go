/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/di-string-match/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/9 09:34
 */

package _di_string_match

func diStringMatch(s string) []int {
	n := len(s)
	ret := make([]int, n+1)

	l, r := 0, n
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ret[i] = l
			l++
		} else {
			ret[i] = r
			r--
		}
	}

	ret[n] = max(l, r)

	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
