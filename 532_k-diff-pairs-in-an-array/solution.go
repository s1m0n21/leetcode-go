/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/k-diff-pairs-in-an-array/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/16 09:55
 */

package _k_diff_pairs_in_an_array

import "sort"

func findPairs(nums []int, k int) int {
	var ret int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left <= right {
			mid := (right-left)/2 + left
			n := nums[mid] - nums[i]
			if n > k {
				right = mid - 1
			} else if n < k {
				left = mid + 1
			} else {
				ret++
				break
			}
		}
	}

	return ret
}
