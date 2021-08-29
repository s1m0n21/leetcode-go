/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/count-of-matches-in-tournament/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/29 3:21 下午
 */

package _count_of_matches_in_tournament

// You can just return (n - 1)
func numberOfMatches(n int) int {
	var match func(n int)
	var matches int

	match = func(n int) {
		if n == 1 {
			return
		}

		if n % 2 != 0 {
			matches++
		}
		n /= 2
		matches += n
		match(n)
	}

	match(n)

	return matches
}