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

	var res []byte
	b := []byte(s)

	for i := 0; i < len(b); i++ {
		l, r := i, i
		m := false
		for {
			if l-1 >= 0 && r+1 < len(b) && b[l-1] == b[r+1] {
				l--
				r++
				m = true
			} else if l-1 >= 0 && b[l-1] == b[i] && !m {
				l--
			} else if r+1 < len(b) && b[r+1] == b[i] && !m {
				r++
			} else {
				break
			}
		}

		if r+1-l > len(res) {
			res = b[l:r+1]
		}
	}

	return string(res)
}
