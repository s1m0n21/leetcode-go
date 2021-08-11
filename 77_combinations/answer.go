/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/combinations/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/11 3:04 下午
 */

package _combinations

func combine(n int, k int) [][]int {
	var traceback func(n, k, st int, c []int) [][]int

	traceback = func(n, k, st int, c []int) [][]int {
		var res [][]int

		if len(c) == k {
			res = append(res, append([]int(nil), c...))
			return  res
		}

		for i := st; i <= n - (k - len(c)) + 1; i++ {
			c = append(c, i)
			res = append(res, traceback(n, k, i+1, c)...)
			c = c[:len(c)-1]
		}

		return res
	}

	return traceback(n, k, 1, []int{})
}
