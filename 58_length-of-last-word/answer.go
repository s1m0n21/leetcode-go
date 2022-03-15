/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/length-of-last-word/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/15 9:05 PM
 */

package _length_of_last_word

import "strings"

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)
	return len(words[len(words)-1])
}
