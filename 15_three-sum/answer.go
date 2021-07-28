/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/3sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/5 12:48 上午
 */

package _three_sum

import  "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	out := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		first := nums[i]
		if first > 0 {
			break
		}
		if i > 0 && first == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			second, third := nums[l], nums[r]
			if first + second + third == 0 {
				out = append(out, []int{first, second, third})
				for l < r && nums[l] == second {
					l++
				}
				for l < r && nums[r] == third {
					r--
				}
			} else if first + second + third < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return out
}
