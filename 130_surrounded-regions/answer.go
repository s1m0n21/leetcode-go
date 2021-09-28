/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/surrounded-regions/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/27 4:38 下午
 */

package _surrounded_regions

var offset = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var m, n int

func dfs(board [][]byte, x, y int) {
	if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != 'O' {
		return
	}

	board[x][y] = 'A'
	for i := 0; i < 4; i++ {
		dfs(board, x+offset[i][0], y+offset[i][1])
	}
}

func solve(board [][]byte) {
	m, n = len(board), len(board[0])
	if m < 3 || n < 3 {
		return
	}

	for i := 0; i < m; i++ {
		if i == 0 || i == m-1 {
			for j := 0; j < n; j++ {
				dfs(board, i, j)
			}
		} else {
			dfs(board, i, 0)
			dfs(board, i, n-1)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case 'A':
				board[i][j] = 'O'
			case 'O':
				board[i][j] = 'X'
			}
		}
	}
}
