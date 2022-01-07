/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/triangle/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/7 3:23 PM
 */

package _triangle

import "math"

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = triangle[0][0]

	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
		}
		dp[i][i] = dp[i-1][i-1] + triangle[i][i]
	}

	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		ans = min(ans, dp[n-1][i])
	}

	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
