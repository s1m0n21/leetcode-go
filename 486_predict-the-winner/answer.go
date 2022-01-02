/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/predict-the-winner/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/29 9:33 AM
 */

package _predict_the_winner

func PredictTheWinner(nums []int) bool {
	var n = len(nums)
	var dp = make([][]int, n)
	var pre = make([]int, n+1)
	var dfs func(l, r int) int

	dfs = func(l, r int) int {
		if dp[l][r] != -1 {
			return dp[l][r]
		}

		if l == r {
			dp[l][r] = nums[r]
		} else if l+1 == r {
			if nums[l] > nums[r] {
				dp[l][r] = nums[l]
			} else {
				dp[l][r] = nums[r]
			}
		} else {
			a, b := dfs(l+1, r), dfs(l, r-1)
			if a < b {
				dp[l][r] = pre[r+1] - pre[l] - a
			} else {
				dp[l][r] = pre[r+1] - pre[l] - b
			}
		}

		return dp[l][r]
	}

	var t = make([]int, n)
	for i := 0; i < n; i++ {
		t[i] = -1
	}

	for i := 0; i < n; i++ {
		dp[i] = append(dp[i], t...)
	}

	for i := 1; i <= n; i++ {
		pre[i] = pre[i-1] + nums[i-1]
	}

	return dfs(0, n-1)*2 >= pre[n]
}
