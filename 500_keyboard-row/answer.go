/*
 * @Author       : s1m0n21
 * @Description  : Answer for https://leetcode-cn.com/problems/keyboard-row/
 * @Date         : 2021/10/31 11:40 PM
 */

package _keyboar_drow

import (
	"strings"
)

var s = struct{}{}
var rows = []map[byte]struct{}{
	{'q': s, 'w': s, 'e': s, 'r': s, 't': s, 'y': s, 'u': s, 'i': s, 'o': s, 'p': s},
	{'a': s, 's': s, 'd': s, 'f': s, 'g': s, 'h': s, 'j': s, 'k': s, 'l': s},
	{'z': s, 'x': s, 'c': s, 'v': s, 'b': s, 'n': s, 'm': s},
}

func findWords(words []string) []string {
	var ans []string
	for _, word := range words {
		str := strings.ToLower(word)
		for _, row := range rows {
			found := true
			for _, w := range str {
				if _, in := row[byte(w)]; !in {
					found = false
					break
				}
			}
			if found {
				ans = append(ans, word)
			}
		}
	}

	return ans
}
