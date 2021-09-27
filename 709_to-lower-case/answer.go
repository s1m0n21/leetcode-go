/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/to-lower-case/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/27 2:16 下午
 */

package _to_lower_case

func toLowerCase(s string) string {
	b := []byte(s)
	for i := 0; i< len(b); i++ {
		if b[i] >= 'A' && b[i] <= 'Z' {
			b[i] += 32
		}
	}
	return string(b)
}
