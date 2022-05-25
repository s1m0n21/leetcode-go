/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/unique-substrings-in-wraparound-string/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/25 17:56
 */

package _unique_substrings_in_wraparound_string

func findSubstringInWraproundString(p string) int {
	var sum int
	var	ch [26]int
	dp := make([]int, len(p))

	for i := 0; i < len(p); i++ {
		if i > 0 && (p[i] - p[i-1] == 1 || (p[i-1] == 'z' && p[i] == 'a')) {
			dp[i] += dp[i-1]
		}
		dp[i]++

		if ch[p[i]-'a'] < dp[i] {
			sum = sum - ch[p[i]-'a'] + dp[i]
			ch[p[i]-'a'] = dp[i]
		}
	}

	return sum
}
