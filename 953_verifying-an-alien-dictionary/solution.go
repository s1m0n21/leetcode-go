/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/verifying-an-alien-dictionary/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/17 09:29
 */

package _verifying_an_alien_dictionary

func isAlienSorted(words []string, order string) bool {
	var ch [26]int
	for i, b := range order {
		ch[b-'a'] = i
	}

	for i := 1; i < len(words); i++ {
		var l, r int
		for l < len(words[i-1]) || r < len(words[i]) {
			if l < len(words[i-1]) && r < len(words[i]) && words[i-1][l] != words[i][r] {
				if ch[words[i-1][l]-'a'] < ch[words[i][r]-'a'] {
					break
				} else {
					return false
				}
			} else if l < len(words[i-1]) && r >= len(words[i]) {
				return false
			}
			l++
			r++
		}
	}

	return true
}
