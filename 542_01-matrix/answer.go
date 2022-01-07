/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/01-matrix/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/6 8:03 PM
 */

package _01_matrix

var directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func updateMatrix(mat [][]int) [][]int {
	var m, n = len(mat), len(mat[0])
	var dist = make([][]int, m)
	var visited = make([][]bool, m)
	var queue [][2]int

	for i := 0; i < m; i++ {
		dist[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
				visited[i][j] = true
			}
		}
	}

	for len(queue) > 0 {
		i, j := queue[0][0], queue[0][1]
		queue = queue[1:]
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && ni < m && nj >= 0 && nj < n && !visited[ni][nj] {
				dist[ni][nj] = dist[i][j] + 1
				queue = append(queue, [2]int{ni, nj})
				visited[ni][nj] = true
			}
		}
	}

	return dist
}
