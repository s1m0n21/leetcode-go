/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/number-of-substrings-with-only-1s/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/12 4:52 下午
 */

package _number_of_substrings_with_only_1s

func numSub(s string) int {
	mod := int(1e9 + 7)
	sum, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			sum = (sum + (count+1)*count/2) % mod
			count = 0
		}
	}
	if count != 0 {
		sum = (sum + (count+1)*count/2) % mod
	}
	return sum
}
