/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/27 9:30 上午
 */

package _median_of_two_sorted_arrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64

	merged := append(nums1, nums2...)
	sort.Ints(merged)
	length := len(merged)

	if length < 2 {
		if len(merged) == 1 {
			res = float64(merged[0])
		}
		return res
	}

	middle := length / 2
	if length % 2 == 0 {
		return float64(merged[middle] + merged[middle-1]) / 2
	} else {
		return float64(merged[middle])
	}
}
