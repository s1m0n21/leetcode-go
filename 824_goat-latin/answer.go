/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/goat-latin/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/28 09:32
 */

package _goat_latin

import "strings"

var vowels = map[byte]struct{}{
	'a': {}, 'A': {},
	'e': {}, 'E': {},
	'i': {}, 'I': {},
	'o': {}, 'O': {},
	'u': {}, 'U': {},
}

func toGoatLatin(sentence string) string {
	var ans []byte

	for i, word := range strings.Fields(sentence) {
		if i > 0 {
			ans = append(ans, ' ')
		}

		temp := append([]byte(nil), word...)

		if _, in := vowels[temp[0]]; in {
			temp = append(temp, 'm', 'a')
		} else {
			temp = append(temp, temp[0], 'm', 'a')
			temp = temp[1:]
		}

		for j := 0; j <= i; j++ {
			temp = append(temp, 'a')
		}

		ans = append(ans, temp...)
	}

	return string(ans)
}
