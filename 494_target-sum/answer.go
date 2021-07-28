/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/target-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/7 1:56 上午
 */

package _target_sum

func findTargetSumWays(nums []int, target int) int {
	var count = 0
	var backtrack func(int, int)

	backtrack = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				count++
			}
			return
		}
		backtrack(idx + 1, sum+nums[idx])
		backtrack(idx + 1, sum-nums[idx])
	}
	backtrack(0, 0)

	return count
}
