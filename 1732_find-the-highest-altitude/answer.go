/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/find-the-highest-altitude/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/1 17:22
 */

package _find_the_highest_altitude

func largestAltitude(gain []int) int {
	var highest int
	for i := 0; i < len(gain); i++ {
		if i > 0 {
			gain[i] += gain[i-1]
		}
		if gain[i] > highest {
			highest = gain[i]
		}
	}

	return highest
}
