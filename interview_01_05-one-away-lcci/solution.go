/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/one-away-lcci/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/13 09:38
 */

package interview_01_05_one_away_lcci

func oneEditAway(first string, second string) bool {
	m, n := len(first), len(second)
	if m < n {
		first, second = second, first
		m, n = n, m
	}

	if m - n > 1 {
		return false
	}

	for i, ch := range second {
		if first[i] != byte(ch) {
			if m == n {
				return first[i+1:] == second[i+1:]
			}
			return first[i+1:] == second[i:]
		}
	}

	return true
}
