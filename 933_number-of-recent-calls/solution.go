/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/number-of-recent-calls/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/6 09:31
 */

package _number_of_recent_calls

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)
	for rc.queue[0] < t-3000 {
		rc.queue = rc.queue[1:]
	}
	return len(rc.queue)
}
