/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/palindrome-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/7 11:19 上午
 */

package _palindrome_number

func isPalindrome(x int) bool {
	var r = 0
	var n = x

	if x < 0 {
		return false
	}

	for x != 0 {
		n := x % 10
		r = r* 10 + n
		if r > 1 << 31 || r < -1 << 31 {
			break
		}
		x /= 10
	}

	return r == n
}
