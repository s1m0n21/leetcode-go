/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/self-dividing-numbers/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/31 13:58
 */

package _self_dividing_numbers

func selfDividingNumbers(left int, right int) []int {
	var ans []int
	for i := left; i <= right; i++ {
		n := i
		skip := false
		for n > 0 {
			d := n % 10
			if d == 0 || i%d != 0 {
				skip = true
				break
			}
			n /= 10
		}
		if !skip {
			ans = append(ans, i)
		}
	}

	return ans
}
