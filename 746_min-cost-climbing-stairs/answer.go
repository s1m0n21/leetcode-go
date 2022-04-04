/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/min-cost-climbing-stairs/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/4 12:00
 */

package _min_cost_climbing_stairs

func minCostClimbingStairs(cost []int) int {
	pre, curr := 0, 0
	for i := 2; i <= len(cost); i++ {
		pre, curr = curr, min(curr+cost[i-1], pre+cost[i-2])
	}
	return curr
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
