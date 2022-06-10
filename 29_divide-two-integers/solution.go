/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/divide-two-integers/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/10 09:34
 */

package _divide_two_integers

func divide(dividend int, divisor int) int {
	neg := dividend^divisor < 0
	dividend = abs(dividend)
	divisor = abs(divisor)
	quotient := 0

	for i := 31; i >= 0; i-- {
		if dividend>>i >= divisor {
			quotient += 1 << i
			dividend -= divisor << i
		}
	}

	if neg {
		quotient *= -1
	}

	if quotient > (1<<31)-1 {
		return (1 << 31) - 1
	}

	return quotient
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return x * -1
}
