/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/non-overlapping-intervals/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/14 2:57 PM
 */

package _non_overlapping_intervals

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans, right := 1, intervals[0][1]
	for _, p := range intervals[1:] {
		if p[0] >= right {
			ans++
			right = p[1]
		}
	}
	return n - ans
}
