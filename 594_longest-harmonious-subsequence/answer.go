/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/longest-harmonious-subsequence/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/20 2:07 下午
 */

package _longest_harmonious_subsequence

import "sort"

func findLHS(nums []int) int {
	sort.Ints(nums)
	n, l, r := 0, 0, 0

	for ; r < len(nums); r++ {
		for ; l < r && nums[r]-nums[l] > 1; l++ {}
		if nums[r]-nums[l] == 1 && r-l+1 > n {
			n = r - l + 1
		}
	}

	return n
}
