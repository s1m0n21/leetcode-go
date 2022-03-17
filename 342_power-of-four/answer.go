/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/power-of-four/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/17 19:39
 */

package _power_of_four

func isPowerOfFour(n int) bool {
	return n > 0 && n&(n-1) == 0 && n&0xaaaaaaaa == 0
}
