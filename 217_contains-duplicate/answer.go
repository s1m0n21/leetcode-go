/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/contains-duplicate/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/17 11:28 下午
 */

package _contains_duplicate

import "sort"

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}
