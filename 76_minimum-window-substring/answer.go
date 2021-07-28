/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/minimum-window-substring/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/30 2:14 下午
 */

package _minimum_window_substring

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var need, curr = 0, 0
	var appeared [58]int
	var target [58]int
	for i := 0; i < len(t); i++ {
		if target[t[i]-'A'] == 0 {
			need++
		}

		target[t[i]-'A']++
	}

	l, r := 0, -1
	out := ""

	for l < len(s) {
		if r+1 < len(s) && need != curr {
			r++
			appeared[s[r]-'A']++
			if appeared[s[r]-'A'] == target[s[r]-'A'] {
				curr++
			}
		} else {
			appeared[s[l]-'A']--
			if appeared[s[l]-'A'] < target[s[l]-'A'] {
				curr--
			}
			l++
		}

		if curr == need {
			if len(out) > (r-l+1) || out == "" {
				out = s[l:r+1]
			}
		}
	}

	return out
}
