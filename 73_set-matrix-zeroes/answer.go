/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/set-matrix-zeroes/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/17 2:26 PM
 */

package _set_matrix_zeroes

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	col0 := false
	for _, r := range matrix {
		if r[0] == 0 {
			col0 = true
		}
		for j := 1; j < n; j++ {
			if r[j] == 0 {
				r[0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if col0 {
			matrix[i][0] = 0
		}
	}
}
