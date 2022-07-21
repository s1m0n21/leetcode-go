/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/shift-2d-grid/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/20 10:08
 */

package _shift_2d_grid

func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
	offset := k % (m * n)
	if offset == 0 {
		return grid
	}

	newGrid := make([][]int, m)
	for i := range newGrid {
		newGrid[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			x, y := i, j
			y -= offset
			for y < 0 {
				x--
				if x < 0 {
					x = m - 1
				}
				y += n
			}
			newGrid[i][j] = grid[x][y]
		}
	}

	return newGrid
}
