/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/maximum-number-of-words-found-in-sentences/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/2 16:00
 */

package _maximum_number_of_words_found_in_sentences

import "strings"

func mostWordsFound(sentences []string) int {
	var ret int
	for _, s := range sentences {
		n := len(strings.Fields(s))
		if n > ret {
			ret = n
		}
	}

	return ret
}
