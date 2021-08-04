/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/permutations-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/4 5:01 下午
 */

package _permutations_ii

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int)
	var used = make([]bool, len(nums))

	sort.Ints(nums)

	backtrack = func(track []int) {
		if len(track) == len(nums) {
			res = append(res, append([]int(nil), track...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if i - 1 >= 0 && nums[i-1] == nums[i] && !used[i-1] {
				continue
			}
			if used[i] {
				continue
			}

			track = append(track, nums[i])
			used[i] = true
			backtrack(track)
			track = track[:len(track)-1]
			used[i] = false
		}
	}

	backtrack(make([]int, 0))

	return res
}
