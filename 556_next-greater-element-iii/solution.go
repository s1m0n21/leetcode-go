/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/next-greater-element-iii/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/4 09:49
 */

package _next_greater_element_iii

import (
	"strconv"
)

func nextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))

	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}

	j := len(nums) - 1
	for j >= 0 && nums[i] >= nums[j] {
		j--
	}

	nums[i], nums[j] = nums[j], nums[i]
	reverse(nums[i+1:])

	ans, _ := strconv.Atoi(string(nums))
	if ans > 1<<31-1 {
		return -1
	}
	return ans
}

func reverse(a []byte) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}
