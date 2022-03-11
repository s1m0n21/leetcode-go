/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/base-7/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/11 5:34 PM
 */

package _base_7

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	var neg bool
	var bytes []byte

	if num < 0 {
		num = -1 * num
		neg = true
	}

	for num > 0 {
		bytes = append(bytes, byte('0'+num%7))
		num /= 7

	}

	if neg {
		bytes = append(bytes, '-')
	}

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}

	return string(bytes)
}
