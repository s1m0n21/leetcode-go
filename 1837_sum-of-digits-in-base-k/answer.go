/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/sum-of-digits-in-base-k/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/14 11:14 AM
 */

package _sum_of_digits_in_base_k

func sumBase(n int, k int) int {
	var sum int

	for n > 0 {
		sum += n % k
		n /= k
	}

	return sum
}
