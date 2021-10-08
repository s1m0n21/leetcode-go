/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/number-of-segments-in-a-string/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/8 9:22 上午
 */

package _number_of_segments_in_a_string

func countSegments(s string) int {
	var res int
	for i, b := range s {
		if (i == 0 || s[i-1] == ' ') && b != ' ' {
			res++
		}
	}
	return res
}
