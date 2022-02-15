/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/longest-palindrome/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/15 11:08 AM
 */

package _longest_palindrome

func longestPalindrome(s string) int {
	var ans int
	m := make(map[rune]int)
	for _, b := range s {
		m[b]++
	}

	for _, cnt := range m {
		ans += cnt / 2 * 2
		if cnt%2 == 1 && ans%2 == 0 {
			ans++
		}
	}

	return ans
}
