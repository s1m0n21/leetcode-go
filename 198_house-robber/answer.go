/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/house-robber/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/2 17:16
 */

package _house_robber

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
