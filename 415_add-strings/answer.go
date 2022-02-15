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

	for p1 >= 0 || p2 >= 0 {
		var sum int
		if p1 >= 0 && p2 >= 0 {
			sum = int(num1[p1]-'0') + int(num2[p2]-'0') + n
			p1--
			p2--
		} else if p1 >= 0 {
			sum = int(num1[p1]-'0') + n
			p1--
		} else {
			sum = int(num2[p2]-'0') + n
			p2--
		}
		n = sum / 10
		sum %= 10
		bytes = append(bytes, byte(sum+'0'))
	}

	if n > 0 {
		bytes = append(bytes, byte(n+'0'))
	}

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}

	return string(bytes)
}
