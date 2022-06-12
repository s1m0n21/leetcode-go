/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/find-and-replace-pattern/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/12 15:30
 */

package _find_and_replace_pattern

func findAndReplacePattern(words []string, pattern string) []string {
	var ret []string

	for _, word := range words {
		if len(word) != len(pattern) {
			continue
		}

		p := make(map[byte]byte)
		w := make(map[byte]byte)
		match := true

		for i := range word {
			if b, in := p[pattern[i]]; in && b != word[i] {
				match = false
				break
			}

			if b, in := w[word[i]]; in && b != pattern[i] {
				match = false
				break
			}

			p[pattern[i]] = word[i]
			w[word[i]] = pattern[i]
		}

		if match {
			ret = append(ret, word)
		}
	}

	return ret
}
