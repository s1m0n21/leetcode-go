/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/14 3:10 下午
 */

package _longest_continuous_increasing_subsequence

func findLengthOfLCIS(nums []int) int {
	l, r, ans := 0, 0, 0
	for ; r < len(nums); r++ {
		if r != 0 && nums[r-1] >= nums[r] {
			l = r
		}

		if r-l+1 > ans {
			ans = r - l + 1
		}
	}
	return ans
}
