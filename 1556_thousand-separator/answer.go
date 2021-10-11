/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/thousand-separator/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/11 3:31 下午
 */

package _thousand_separator

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	var b []byte
	var i int
	for n > 0 {
		if i != 0 && i % 3 == 0 {
			b = append(b, '.')
		}
		m := n % 10
		n /= 10
		b = append(b, byte('0'+m))
		i++
	}

	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	return string(b)
}
