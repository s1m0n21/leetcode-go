/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/minimum-size-subarray-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/29 1:58 下午
 */

package _minimum_size_subarray_sum

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, -1
	out := len(nums) + 1
	sum := 0

	for left < len(nums) {
		if sum < target && right+1 < len(nums) {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= target {
			out = min(out, right-left+1)
		}
	}

	if out == len(nums)+1 {
		return 0
	}

	return out
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
