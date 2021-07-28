/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/happy-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/7/2 1:58 上午
 */

package _happy_number

func isHappy(n int) bool {
	var appeared = make(map[int]int)

	for n != 1 {
		curr := 0
		for {
			if n < 10 {
				curr += n * n
				break
			}
			curr += (n % 10) * (n % 10)
			n /= 10
		}

		appeared[curr]++
		if appeared[curr] > 1 {
			return false
		}

		n = curr
	}

	return true
}
