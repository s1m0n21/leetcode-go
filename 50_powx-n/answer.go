/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/powx-n/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/8 11:43 下午
 */

package _powx_n

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n < 0 {
		x = 1 / x
		n = -n
	}


	t := myPow(x, n/2)
	r := t * t
	if n % 2 == 1 {
		r = r * x
	}

	return r
}
