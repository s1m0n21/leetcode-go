/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/teemo-attacking/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/10 10:54 上午
 */

package _teemo_attacking

func findPoisonedDuration(timeSeries []int, duration int) int {
	n := duration * len(timeSeries)
	for i := 1; i < len(timeSeries); i++ {
		if timeSeries[i] - timeSeries[i-1] < duration {
			n -= duration - (timeSeries[i] - timeSeries[i-1])
		}
	}
	return n
}
