/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/climbing-stairs/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/12 3:04 下午
 */

package _climbing_stairs

func climbStairs(n int) int {
	memo := make([]int, n+1)

	memo[0] = 1
	memo[1] = 1
	for i := 2; i <= n ; i++ {
		memo[i] = memo[i-1] + memo[i-2]
	}

	return memo[n]
}
