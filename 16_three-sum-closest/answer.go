/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/3sum-closest/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/5 5:44 下午
 */

package _three_sum_closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var closest = math.MaxInt32

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			}

			if abs(target - sum) < abs(target - closest) {
				closest = sum
			}

			if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return closest
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
