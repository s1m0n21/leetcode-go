/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/corporate-flight-bookings/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/8/31 12:24 上午
 */

package _corporate_flight_bookings

func corpFlightBookings(bookings [][]int, n int) []int {
	res := make([]int, n)

	for _, booking := range bookings {
		res[booking[0]-1] += booking[2]
		if booking[1] < n {
			res[booking[1]] -= booking[2]
		}
	}

	for i := 1; i < n; i++ {
		res[i] += res[i-1]
	}

	return res
}
