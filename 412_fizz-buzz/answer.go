/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/fizz-buzz/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/13 9:37 上午
 */

package _fizz_buzz

import "strconv"

func fizzBuzz(n int) []string {
	var ans = make([]string, n)
	for i := 1; i <= n; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i % 3 == 0 {
			ans[i-1] = "Fizz"
		} else if i % 5 == 0 {
			ans[i-1] = "Buzz"
		} else {
			ans[i-1] = strconv.Itoa(i)
		}
	}
	return ans
}
