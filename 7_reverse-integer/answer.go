/**
 * @Author         : simon
 * @Description    : Answer for https://leetcode-cn.com/problems/reverse-integer/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/4 10:38 上午
 */

package _reverse_integer

func reverse(x int) int {
	var out = 0

	for x != 0 {
		n := x % 10
		out = out * 10 + n

		if out > 1 << 31 || out < -1 << 31 {
			return 0
		}

		x /= 10
	}

	return out
}
