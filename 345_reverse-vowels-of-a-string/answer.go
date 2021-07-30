/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/29 2:18 上午
 */

package _reverse_vowels_of_a_string

var vowels = map[byte]struct{}{
	'a': {}, 'A': {},
	'e': {}, 'E': {},
	'i': {}, 'I': {},
	'o': {}, 'O': {},
	'u': {}, 'U': {},
}

func reverseVowels(s string) string {
	bytes := []byte(s)
	i, j := 0, len(s)-1
	for i <= j {
		if !isVowel(s[i]) {
			i++
		} else if !isVowel(s[j]) {
			j--
		} else {
			bytes[i] = s[j]
			bytes[j] = s[i]

			i++
			j--
		}
	}

	return string(bytes)
}

func isVowel(s byte) bool {
	if _, has := vowels[s]; has {
		return true
	}
	return false
}
