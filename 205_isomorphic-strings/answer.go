/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/isomorphic-strings/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/4 12:13 上午
 */

package _isomorphic_strings

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sindex [256]byte
	var tindex [256]byte
	for i := 0; i < len(s); i++ {
		if sindex[s[i]-'a'] > 0 && sindex[s[i]-'a'] != t[i] || tindex[t[i]-'a'] > 0 && tindex[t[i]-'a'] != s[i] {
			return false
		}

		sindex[s[i]-'a'] = t[i]
		tindex[t[i]-'a'] = s[i]
	}

	return true
}
