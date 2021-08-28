/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/running-sum-of-1d-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/28 10:08 下午
 */

package _running_sum_of_1d_array

func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
