/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/combination-sum-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/12 10:16 上午
 */

package _combination_sum_ii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)

	var traceback func(c []int, st, sum int) [][]int

	traceback = func(c []int, st, sum int) [][]int {
		var res [][]int
		if sum >= target {
			if sum == target {
				res = append(res, append([]int(nil), c...))
			}
			return res
		}

		for i := st; i < len(candidates); i++ {
			if i-1 >= st && candidates[i] == candidates[i-1] {
				continue
			}

			c = append(c, candidates[i])
			res = append(res, traceback(c, i+1, sum+candidates[i])...)
			c = c[:len(c)-1]
		}

		return res
	}

	return traceback([]int{}, 0, 0)
}
