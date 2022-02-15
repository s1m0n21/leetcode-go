/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/increasing-triplet-subsequence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/15 11:27 PM
 */

package _increasing_triplet_subsequence

func increasingTriplet(nums []int) bool {
	var first, second = nums[0], 1<<31 - 1
	for _, num := range nums {
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}
