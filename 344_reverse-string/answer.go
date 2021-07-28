/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/29 2:11 上午
 */

package _reverse_string

func reverseString(s []byte)  {
	i, j := 0, len(s)-1
	for i < j {
		b := s[i]
		s[i] = s[j]
		s[j] = b
		i++
		j--
	}
}
