/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/string-to-integer-atoi/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/5 2:05 上午
 */

package _string_to_integer_atoi

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	var sign int
	var abs string

	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}

	switch str[0] {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		sign, abs = 1, str
	case '+':
		sign, abs = 1, str[1:]
	case '-':
		sign, abs = -1, str[1:]
	default:
		abs = ""
		return 0
	}

	for i, b := range abs {
		if b < '0' || b > '9' {
			abs = abs[:i]
			break
		}
	}

	num := 0
	for _, b := range abs {
		num = num*10 + int(b-'0')
		if sign == 1 && num > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign == -1 && sign*num < math.MinInt32 {
			return math.MinInt32
		}
	}

	return sign * num
}
