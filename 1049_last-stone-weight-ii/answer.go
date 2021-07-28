/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/last-stone-weight-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/8 10:43 上午
 */

package _last_stone_weight_ii

func lastStoneWeightII(stones []int) int {
	var sum = 0
	for _, w := range stones {
		sum += w
	}
	m := sum / 2
	dp := make([]bool, m+1)
	dp[0] = true
	for _, w := range stones {
		for j := m; j>= w; j-- {
			dp[j] = dp[j] || dp[j-w]
		}
	}
	for j := m; ; j-- {
		if dp[j] {
			return sum - 2 * j
		}
	}
}
