/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/search-a-2d-matrix-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/13 9:59 PM
 */

package _search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		l, r := 0, n-1
		for l <= r {
			mid := l + (r-l)/2
			if matrix[i][mid] < target {
				l = mid + 1
			} else if matrix[i][mid] > target {
				r = mid - 1
			} else {
				return true
			}
		}
	}

	return false
}
