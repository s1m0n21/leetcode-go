/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/max-area-of-island/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/6 12:34 PM
 */

package _max_area_of_island

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func maxAreaOfIsland(grid [][]int) int {
	var ans int
	var m, n = len(grid), len(grid[0])
	var dfs func(x, y int) int

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dfs = func(x, y int) int {
		if x < 0 || x >= m || y < 0 || y >= n || visited[x][y] {
			return 0
		}

		var cnt int
		visited[x][y] = true

		if grid[x][y] == 1 {
			cnt++
			for _, d := range directions {
				cnt += dfs(x+d[0], y+d[1])
			}
		}

		return cnt
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !visited[i][j] {
				cnt := dfs(i, j)
				if cnt > ans {
					ans = cnt
				}
			}
		}
	}

	return ans
}
