/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/climbing-stairs/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/12 3:04 下午
 */

package _climbing_stairs

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	var a, b, t = 1, 2, 0
	for i := 3; i <= n; i++ {
		t = a + b
		a = b
		b = t
	}

	return t
}
