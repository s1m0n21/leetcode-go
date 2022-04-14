/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/richest-customer-wealth/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/14 12:57
 */

package _richest_customer_wealth

func maximumWealth(accounts [][]int) int {
	ret := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > ret {
			ret = sum
		}
	}
	return ret
}
