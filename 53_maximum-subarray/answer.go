/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/maximum-subarray/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/11 7:49 PM
 */

package _maximum_subarray

func maxSubArray(nums []int) int {
	var max = nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
