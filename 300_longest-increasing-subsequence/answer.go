/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/longest-increasing-subsequence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/6 7:47 PM
 */

package _longest_increasing_subsequence

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var ans int
	var dp = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > ans {
			ans = dp[i]
		}
	}

	return ans
}
