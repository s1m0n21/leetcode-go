/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-words-in-a-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/1 2:46 下午
 */

package _reverse_words_in_a_string

import "strings"

func reverseWords(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-i-1] = words[len(words)-i-1], words[i]
	}

	res := ""
	for i, word := range words {
		if i == 0 {
			res += word
		} else {
			res += " " + word
		}
	}

	return res
}
