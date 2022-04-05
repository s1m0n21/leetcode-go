/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/squares-of-a-sorted-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/5 13:56
 */

package _squares_of_a_sorted_array

import "sort"

func sortedSquares(nums []int) []int {
	for i, n := range nums {
		nums[i] = n * n
	}
	sort.Ints(nums)
	return nums
}
