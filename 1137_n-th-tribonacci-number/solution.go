/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/n-th-tribonacci-number/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/14 15:00
 */

package _n_th_tribonacci_number

func tribonacci(n int) int {
	if n < 3 {
		if n == 0 {
			return 0
		}
		return 1
	}

	t1, t2, t3 := 0, 1, 1
	for i := 3; i <= n; i++ {
		t1, t2, t3 = t2, t3, t1+t2+t3
	}

	return t3
}
