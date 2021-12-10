/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/find-peak-element/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/15 9:43 上午
 */

package _find_peak_element

import "math"

func findPeakElement(nums []int) int {
	get := func(i int) int {
		if i < 0 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	l, r := 0, len(nums)-1
	for {
		mid := (l + r) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
}
