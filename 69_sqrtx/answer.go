/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/sqrtx/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/15 9:37 PM
 */

package _sqrtx

func mySqrt(x int) int {
	n := x
	for n*n > x {
		n = (n + x/n) / 2
	}
	return n
}
