/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/generate-parentheses/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/13 11:09 上午
 */

package _generate_parentheses

func generateParenthesis(n int) []string {
	var traceback func(l, r int, p string)
	var res []string

	traceback = func(l, r int, p string) {
		if len(p) == n*2 {
			res = append(res, p)
			return
		}

		if l > 0 {
			traceback(l-1, r, p+"(")
		}
		if l < r {
			traceback(l, r-1, p+")")
		}
	}

	traceback(n, n, "")

	return res
}
