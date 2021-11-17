/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/distance-between-bus-stops/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/17 3:36 下午
 */

package _distance_between_bus_stops

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var d1, d2 = 0, 0
	n := len(distance)

	curr := start
	for curr != destination {
		d1 += distance[curr]
		curr = (curr + 1) % n
	}

	curr = destination
	for curr != start {
		d2 += distance[curr]
		curr = (curr + 1) % n
	}

	if d1 < d2 {
		return d1
	}
	return d2
}
