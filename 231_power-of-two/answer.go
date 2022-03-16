/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/power-of-two/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/17 01:34
 */

package _power_of_two

const max = 1 << 30

func isPowerOfTwo(n int) bool {
	return n > 0 && max%n == 0
}
