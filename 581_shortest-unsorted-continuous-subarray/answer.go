/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/3 11:48 下午
 */

package _shortest_unsorted_continuous_subarray

import "sort"

func findUnsortedSubarray(nums []int) int {
	sorted := append([]int(nil), nums...)
	sort.Ints(sorted)

	l, r := 0, len(nums)-1
	for l < r {
		if nums[l] == sorted[l] {
			l++
		} else if nums[r] == sorted[r] {
			r--
		}

		if (nums[l] != sorted[l]) && (nums[r] != sorted[r]) {
			r +=1
			break
		}
	}

	return r - l
}
