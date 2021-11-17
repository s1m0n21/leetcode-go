/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-prefix-of-word/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/17 4:27 下午
 */

package _reverse_prefix_of_word

func reversePrefix(word string, ch byte) string {
	b := []byte(word)
	for i := range b {
		if b[i] == ch {
			for j := 0; j < (i+1)/2; j++ {
				b[j], b[i-j] = b[i-j], b[j]
			}
			break
		}
	}
	return string(b)
}
