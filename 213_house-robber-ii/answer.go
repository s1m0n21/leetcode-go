/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/house-robber-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/7 5:02 PM
 */

package _house_robber_ii

func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(dp(nums[:n-1]), dp(nums[1:]))
}

func dp(nums []int) int {
	a, b := nums[0], max(nums[0], nums[1])
	for _, n := range nums[2:] {
		a, b = b, max(a+n, b)
	}
	return b
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
