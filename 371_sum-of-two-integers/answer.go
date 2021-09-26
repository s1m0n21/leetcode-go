/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sum-of-two-integers/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/26 9:18 上午
 */

package _sum_of_two_integers

func getSum(a int, b int) int {
	if a&b == 0 {
		return a | b
	}
	return getSum(a^b, (a&b)<<1)
}
