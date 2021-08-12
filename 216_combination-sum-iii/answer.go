/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/combination-sum-iii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/12 12:44 下午
 */

package _combination_sum_iii

func combinationSum3(k int, n int) [][]int {
	var traceback func(c []int, st, sum int) [][]int

	traceback = func(c []int, st, sum int) [][]int {
		var res [][]int
		if sum >= n || len(c) == k {
			if sum == n && len(c) == k {
				res = append(res, append([]int(nil), c...))
			}
			return res
		}

		for i := st; i <= 9; i++ {
			c = append(c, i)
			res = append(res, traceback(c, i+1, sum+i)...)
			c = c[:len(c)-1]
		}

		return res
	}

	return traceback([]int{}, 1, 0)
}
