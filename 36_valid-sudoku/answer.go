/**
 * @Author         : s1m0n21
 * @Description    : Answer of valid-sudoku
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/1/17 4:00 PM
 */

package _valid_sudoku

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' && !isValid(board, i, j, board[i][j]) {
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, x, y int, n byte) bool {
	for i := 0; i < 9; i++ {
		if (i != y && board[x][i] == n) || (i != x && board[i][y] == n) {
			return false
		}
	}

	startRow := (x / 3) * 3
	startCol := (y / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if (i != x || j != y) && board[i][j] == n {
				return false
			}
		}
	}

	return true
}
