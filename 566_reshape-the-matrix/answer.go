/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/reshape-the-matrix/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/13 12:39 PM
 */

package _reshape_the_matrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	m, n := len(mat), len(mat[0])
	if m*n != r*c {
		return mat
	}

	ans := make([][]int, r)
	for i := range ans {
		ans[i] = make([]int, c)
	}

	x, y := 0, 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans[x][y] = mat[i][j]
			if y+1 < c {
				y++
			} else {
				x++
				y = 0
			}
		}
	}

	return ans
}
