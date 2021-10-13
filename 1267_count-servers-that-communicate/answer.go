/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/count-servers-that-communicate/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/13 2:34 下午
 */

package _count_servers_that_communicate

func countServers(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	x, y := make([]int, m), make([]int, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				x[i]++
				y[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && (x[i] > 1 || y[j] > 1) {
				ans++
			}
		}
	}

	return ans
}
