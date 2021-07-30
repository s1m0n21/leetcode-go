/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/is-subsequence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/31 1:59 上午
 */

package _is_subsequence

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) > len(t) {
		return false
	}

	k, j := 0, 0
	for k < len(s) && j < len(t) {
		if s[k] == t[j] {
			k++
		}
		j++
	}

	return k == len(s)
}
