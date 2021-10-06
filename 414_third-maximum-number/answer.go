/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/third-maximum-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/6 9:44 下午
 */

package _third_maximum_number

func thirdMax(nums []int) int {
	min := -1<<31 - 1
	a, b, c := min, min, min
	for _, n := range nums {
		if n > a {
			a, b, c = n, a, b
		} else if n > b && n < a {
			b, c = n, b
		} else if n > c && n < b {
			c = n
		}
	}

	if c == min {
		return a
	}

	return c
}
