/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/longest-common-subsequence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/8 4:03 AM
 */

package _longest_common_subsequence

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i, c1 := range text1 {
		for j, c2 := range text2 {
			if c1 == c2 {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
