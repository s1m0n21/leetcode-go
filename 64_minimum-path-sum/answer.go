/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/minimum-path-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/7 4:17 PM
 */

package _minimum_path_sum

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < n; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
