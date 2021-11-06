/*
 * @Author       : s1m0n21
 * @Description  : Answer for https://leetcode-cn.com/problems/sudoku-solver/
 * @Date         : 2021/11/05 4:31 PM
 */

package _sudoku_solver

func solveSudoku(board [][]byte) {
	var backtrack func() bool

	backtrack = func() bool {
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' {
					continue
				}

				var k byte = '1'
				for ; k <= '9'; k++ {
					if isValid(board, i, j, k) {
						board[i][j] = k
						if backtrack() {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}

		return true
	}

	backtrack()
}

func isValid(board [][]byte, row, col int, n byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == n || board[i][col] == n {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if board[i][j] == n {
				return false
			}
		}
	}

	return true
}
