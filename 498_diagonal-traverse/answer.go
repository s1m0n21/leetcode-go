/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/diagonal-traverse/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/10/14 5:11 下午
 */

package _diagonal_traverse

var d = [2][3][2]int{
	{{-1, 1}, {0, 1}, {1, 0}},
	{{1, -1}, {1, 0}, {0, 1}},
}

func findDiagonalOrder(mat [][]int) []int {
	var ans []int
	next := d[0]
	m, n := len(mat), len(mat[0])
	x, y := 0, 0

	isValid := func(x, y int) bool {
		if x < 0 || x >= m || y < 0 || y >= n {
			return false
		}
		return true
	}

	for {
		ans = append(ans, mat[x][y])
		if x == m-1 && y == n-1 {
			break
		}

		if !isValid(x+next[0][0], y+next[0][1]) {
			if isValid(x+next[1][0], y+next[1][1]) {
				x, y = x+next[1][0], y+next[1][1]
			} else {
				x, y = x+next[2][0], y+next[2][1]
			}
			if next == d[0] {
				next = d[1]
			} else {
				next = d[0]
			}
		} else {
			x, y = x+next[0][0], y+next[0][1]
		}
	}

	return ans
}
