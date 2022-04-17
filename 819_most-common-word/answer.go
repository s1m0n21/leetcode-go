/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/most-common-word/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/17 13:07
 */

package _most_common_word

func mostCommonWord(paragraph string, banned []string) string {
	ban := make(map[string]struct{})
	for _, s := range banned {
		ban[s] = struct{}{}
	}

	words := []byte(paragraph)
	count := make(map[string]int)
	l, r := 0, 0
	for r < len(words) {
		b := words[r]
		if b < 'a' || b > 'z' {
			if b+32 >= 'a' && b+32 <= 'z' {
				words[r] = b + 32
			} else {
				s := string(words[l:r])
				if _, has := ban[s]; !has && s != "" {
					count[s]++
				}
				r++
				l = r
				continue
			}
		}
		r++
	}
	s := string(words[l:r])
	if _, has := ban[s]; !has && s != "" {
		count[s]++
	}

	ret := ""
	max := 0
	for word, cnt := range count {
		if cnt > max {
			ret = word
			max = cnt
		}
	}

	return ret
}
