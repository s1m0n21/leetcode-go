/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/sum-of-digits-of-string-after-convert/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/18 4:03 下午
 */

package _sum_of_digits_of_string_after_convert

func getLucky(s string, k int) int {
	var ans int
	for _, b := range s {
		b -= 'a' - 1
		ans += int(b/10 + b%10)
	}
	k--
	for ; k > 0 && ans > 9; k-- {
		n := ans
		for ans = 0; n > 0; n /= 10 {
			ans += n % 10
		}
	}
	return ans
}
