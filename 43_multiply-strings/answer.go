/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/multiply-strings/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/17 7:42 PM
 */

package _multiply_strings

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	var ans = "0"
	m, n := len(num1), len(num2)
	for i := n - 1; i >= 0; i-- {
		var curr []byte
		var add int
		for j := n - 1; j > i; j-- {
			curr = append(curr, '0')
		}
		y := int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			x := int(num1[j] - '0')
			p := x*y + add
			curr = append(curr, byte(p%10+'0'))
			add = p / 10
		}
		for ; add != 0; add /= 10 {
			curr = append(curr, byte(add%10+'0'))
		}
		reverseBytes(curr)
		ans = addString(ans, string(curr))
	}

	return ans
}

func addString(num1 string, num2 string) string {
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
	reverseBytes(bytes)

	return string(bytes)
}

func reverseBytes(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
}
