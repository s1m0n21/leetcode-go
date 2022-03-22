/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/fibonacci-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/22 16:57
 */

package _fibonacci_number

func fib(n int) int {
	if n < 1 {
		return n
	}

	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}

	return b
}
