/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/minimum-falling-path-sum/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/4/3 13:29
 */

package _minimum_falling_path_sum

import (
	"math"
)

var d = [3][2]int{{1, 0}, {1, -1}, {1, 1}}

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			min := math.MaxInt
			for k := 0; k < 3; k++ {
				x, y := i+d[k][0], j+d[k][1]
				if x >= 0 && y >= 0 && x < n && y < n && matrix[i][j]+matrix[x][y] < min {
					min = matrix[i][j] + matrix[x][y]
				}
			}
			matrix[i][j] = min
		}
	}

	min := matrix[0][0]
	for i := 0; i < n; i++ {
		if matrix[0][i] < min {
			min = matrix[0][i]
		}
	}

	return min
}
