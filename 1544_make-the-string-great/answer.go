/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/make-the-string-great/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/29 14:59
 */

package _make_the_string_great

func makeGood(s string) string {
	b := []byte(s)
	for {

		changed := false
		for i := 1; i < len(b); i++ {
			if b[i]^b[i-1] == 32 {
				changed = true
				b = append(append([]byte(nil), b[:i-1]...), b[i+1:]...)
				break
			}
		}

		if !changed {
			break
		}
	}

	return string(b)
}
