/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/number-of-islands/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/27 3:29 下午
 */

package _number_of_islands

var offset = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var m, n int
var visited [][]bool

func numIslands(grid [][]byte) int {
	m, n = len(grid), len(grid[0])
	visited = nil
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nextx, nexty := x + offset[i][0], y + offset[i][1]
			if (nextx >= 0 && nextx < m) && (nexty >= 0 && nexty < n) &&
				(!visited[nextx][nexty]) &&
				(grid[nextx][nexty]) == '1' {
				dfs(nextx, nexty)
			}
		}
	}

	var res int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}
