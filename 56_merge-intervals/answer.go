/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/merge-intervals/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/11 10:26 AM
 */

package _merge_intervals

import "sort"

func merge(intervals [][]int) [][]int {
	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] <= end {
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			ans = append(ans, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}
	ans = append(ans, []int{start, end})

	return ans
}
