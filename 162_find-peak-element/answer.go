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
	idx := len(nums) / 2

	get := func(i int) int {
		if i < 0 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}

	for !(get(idx-1) < get(idx) && get(idx) > get(idx+1)) {
		if get(idx+1) > get(idx) {
			idx++
		} else {
			idx--
		}
	}

	return idx
}
