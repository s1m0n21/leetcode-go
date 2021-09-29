/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/pacific-atlantic-water-flow/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/29 2:08 下午
 */

package _pacific_atlantic_water_flow

var offset = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return nil
	}

	var dfs func(visited [][]bool, last, x, y int)
	var p, a [][]bool
	m, n := len(heights), len(heights[0])
	for i := 0; i < m; i++ {
		p = append(p, make([]bool, n))
		a = append(a, make([]bool, n))
	}

	dfs = func(visited [][]bool, last, x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || heights[x][y] < last || visited[x][y] {
			return
		}

		visited[x][y] = true
		for i := 0; i < 4; i++ {
			dfs(visited, heights[x][y], x+offset[i][0], y+offset[i][1])
		}
	}

	for i := 0; i < m; i++ {
		dfs(p, heights[i][0], i, 0)
		dfs(a, heights[i][n-1], i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(p, heights[0][i], 0, i)
		dfs(a, heights[m-1][i], m-1, i)
	}

	var res [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	return res
}
