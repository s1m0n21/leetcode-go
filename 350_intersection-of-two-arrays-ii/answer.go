/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/30 11:43 下午
 */

package _intersection_of_two_arrays_ii

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	var m, n = len(nums1), len(nums2)

	sort.Ints(nums1)
	sort.Ints(nums2)

	var i, j int
	for i < m && j < n {
		if nums1[i] == nums2[j] {
			ans = append(ans, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			for ; i < m && nums1[i] < nums2[j]; i++ {
			}
		} else if nums2[j] < nums1[i] {
			for ; j < n && nums2[j] < nums1[i]; j++ {
			}
		}
	}

	return ans
}
