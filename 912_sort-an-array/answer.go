/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sort-an-array/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/18 10:00 上午
 */

package _sort_an_array

import "math/rand"

func sortArray(nums []int) []int {
	pivot(nums, 0, len(nums)-1)
	return nums
}

func pivot(nums []int, l, r int) {
	if l >= r {
		return
	}

	mid := rand.Int()%(r-l+1) + l
	nums[l], nums[mid] = nums[mid], nums[l]
	i, j := l+1, r
	x := nums[l]

	for i < j {
		for i < j && nums[i] <= x {
			i++
		}
		for i < j && nums[j] >= x {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	if nums[i] <= x {
		nums[i], nums[l] = nums[l], nums[i]
	} else {
		i--
		nums[i], nums[l] = nums[l], nums[i]
	}

	pivot(nums, i+1, r)
	pivot(nums, l, i-1)
}
