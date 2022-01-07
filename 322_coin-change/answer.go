/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/coin-change/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/7 12:41 PM
 */

package _coin_change

func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				if dp[i-coins[j]]+1 < dp[i] {
					dp[i] = dp[i-coins[j]] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
