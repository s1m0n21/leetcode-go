/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/word-pattern/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/2 2:48 上午
 */

package _word_pattern

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(words) != len(pattern) {
		return false
	}

	var pvalid = [26]string{}
	var svalid = make(map[string]byte)

	for i, word := range words {
		if pvalid[pattern[i]-'a'] != "" && pvalid[pattern[i]-'a'] != word || svalid[word] != 0 && svalid[word] != pattern[i] {
			return false
		}
		pvalid[pattern[i]-'a'] = word
		svalid[word] = pattern[i]
	}

	return true
}
