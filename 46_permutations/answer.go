/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/permutations/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/7 4:34 下午
 */

package _permutations

func permute(nums []int) [][]int {
	var out [][]int
	var backtrack func([]int)
	var index = make(map[int]bool)

	backtrack = func(track []int) {
		if len(track) == len(nums) {
			out = append(out, track)
			return
		}

		for i := 0; i < len(nums); i++ {
			if exist, has :=  index[nums[i]]; exist && has {
				continue
			}

			index[nums[i]] = true
			backtrack(append(track, nums[i]))
			index[nums[i]] = false
		}
	}

	backtrack(make([]int, 0))

	return out
}
