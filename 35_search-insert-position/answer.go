/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/search-insert-position/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/15 9:23 PM
 */

package _search_insert_position

func searchInsert(nums []int, target int) int {
	l, r, ans := 0, len(nums)-1, len(nums)
	for l <= r {
		m := (r-l)/2 + l
		if target <= nums[m] {
			ans = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return ans
}
