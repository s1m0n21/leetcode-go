/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/2 16:24
 */

package _replace_all_s_to_avoid_consecutive_repeating_characters

func modifyString(s string) string {
	b := []byte(s)
	n := len(b)
	for i := 0; i < n; i++ {
		if b[i] == '?' {
			target := byte('a')
			for j := 0; j < 26; j++ {
				if (i == 0 || b[i-1] != target) && (i+1 == n || b[i+1] != target) {
					break
				}
				target++
			}
			b[i] = target
		}
	}

	return string(b)
}
