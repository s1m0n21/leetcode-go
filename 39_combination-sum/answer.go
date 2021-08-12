/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/combination-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/11 9:18 下午
 */

package _combination_sum

func combinationSum(candidates []int, target int) [][]int {
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
			c = append(c, candidates[i])
			res = append(res, traceback(c, i, sum+candidates[i])...)
			c = c[:len(c)-1]
		}

		return res
	}

	return traceback([]int{}, 0, 0)
}
