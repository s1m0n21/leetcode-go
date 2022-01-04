/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/find-right-interval/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/4 12:10 PM
 */

package _find_right_interval

import (
	"sort"
)

func findRightInterval(intervals [][]int) []int {
	var ans = make([]int, len(intervals))
	var idx = make([]int, len(intervals))
	for i := range idx {
		idx[i] = i
	}

	sort.Slice(idx, func(i, j int) bool {
		return intervals[idx[i]][0] < intervals[idx[j]][0]
	})

	for i := range idx {
		if intervals[idx[i]][0] == intervals[idx[i]][1] {
			ans[idx[i]] = i
			continue
		}

		ans[idx[i]] = -1

		l, r := i, len(idx)
		for l < r {
			mid := (l + r) / 2
			if intervals[idx[mid]][0] >= intervals[idx[i]][1] {
				r = mid
			} else {
				l = mid + 1
			}
		}

		if l > i && l < len(idx) {
			ans[idx[i]] = idx[l]
		}
	}

	return ans
}
