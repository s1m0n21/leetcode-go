/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/recursive-mulitply-lcci/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/11 4:14 下午
 */

package interview_08_05_recursive_mulitply_lcci

func multiply(A int, B int) int {
	if B < A {
		A, B = B, A
	}

	var res int
	var mul func(i int)

	mul = func(i int) {
		if i == A {
			return
		}
		res += B
		mul(i+1)
	}
	mul(0)

	return res
}
