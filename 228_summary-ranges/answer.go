/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/summary-ranges/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/1 10:26
 */

package _summary_ranges

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	var ret []string
	start, curr := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i]-curr != 1 {
			if start == curr {
				ret = append(ret, fmt.Sprintf("%d", start))
			} else {
				ret = append(ret, fmt.Sprintf("%d->%d", start, curr))
			}
			start = nums[i]
			curr = nums[i]
			continue
		}
		curr = nums[i]
	}

	if start == curr {
		ret = append(ret, fmt.Sprintf("%d", start))
	} else {
		ret = append(ret, fmt.Sprintf("%d->%d", start, curr))
	}

	return ret
}
