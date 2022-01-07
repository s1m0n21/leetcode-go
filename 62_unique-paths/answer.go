/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/unique-paths/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/6 8:20 PM
 */

package _unique_paths

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}
