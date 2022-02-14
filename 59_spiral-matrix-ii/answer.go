/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/spiral-matrix-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/14 9:42 AM
 */

package _spiral_matrix_ii

var d = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	x, y, k := 0, 0, 0
	for i := 1; i <= n*n; i++ {
		matrix[x][y] = i
		newx, newy := x+d[k%4][0], y+d[k%4][1]
		if newx < 0 || newx >= n || newy < 0 || newy >= n || matrix[newx][newy] != 0 {
			k++
			newx, newy = x+d[k%4][0], y+d[k%4][1]
		}
		x, y = newx, newy
	}

	return matrix
}
