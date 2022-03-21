/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/ugly-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/21 17:53
 */

package _ugly_number

var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}

	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}

	return n == 1
}
