/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/4sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/5 1:06 上午
 */

package _four_sum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	out := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		first := nums[i]
		if i > 0 && nums[i-1] == first {
			continue
		}

		for j := i+1; j < len(nums); j++ {
			second := nums[j]
			if j > i+1 && nums[j-1] == second {
				continue
			}

			l, r := j+1, len(nums)-1
			for l < r {
				third, fourth := nums[l], nums[r]
				sum := first + second + third + fourth
				if sum == target {
					out = append(out, []int{first, second, third, fourth})
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return out
}
