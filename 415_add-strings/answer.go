/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/add-strings/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/15 9:54 AM
 */

package _add_strings

func addStrings(num1 string, num2 string) string {
	var bytes []byte
	var p1, p2, n = len(num1) - 1, len(num2) - 1, 0

	for p1 >= 0 || p2 >= 0 || n > 0 {
		var sum, x, y int
		if p1 >= 0 {
			x = int(num1[p1] - '0')
			p1--
		}
		if p2 >= 0 {
			y = int(num2[p2] - '0')
			p2--
		}
		sum = x + y + n
		n = sum / 10
		bytes = append(bytes, byte(sum%10+'0'))
	}

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}

	return string(bytes)
}
