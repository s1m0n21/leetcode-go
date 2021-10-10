/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/arranging-coins/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/10 10:26 下午
 */

package _arranging_coins

func arrangeCoins(n int) int {
	i := 0
	for n - (i+1) >= 0{
		i++
		n -= i
	}
	return i
}
