/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/longest-palindromic-substring/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/6 3:42 下午
 */

package _longest_palindromic_substring

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	var res string

	for i := 0; i < len(s); i++ {
		l, r := i, i
		m := false
		for {
			if l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
				l--
				r++
				m = true
			} else if l-1 >= 0 && s[l-1] == s[i] && !m {
				l--
			} else if r+1 < len(s) && s[r+1] == s[i] && !m {
				r++
			} else {
				break
			}
		}

		if r+1-l > len(res) {
			res = s[l:r+1]
		}
	}

	return res
}
