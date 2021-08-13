/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/n-queens/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/6/7 9:57 下午
 */

package _n_queeens

func solveNQueens(n int) [][]string {
	var out [][]string
	var backtrack func([]string, int)

	backtrack = func(board []string, row int) {
		if row == len(board) {
			out = append(out, copyBoard(board))
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}

			board[row] = board[row][:col] + "Q" + board[row][col+1:]
			backtrack(board, row + 1)
			board[row] = board[row][:col] + "." + board[row][col+1:]
		}
	}

	var board []string
	for i := 0; i < n; i++ {
		var s = ""
		for j := 0; j < n; j++ {
			s = s + "."
		}
		board = append(board, s)
	}

	backtrack(board, 0)

	return out
}

func copyBoard(board []string) []string {
	var out []string
	for _, row := range board {
		out = append(out, row)
	}

	return out
}

func isValid(board []string, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if string(board[i][col]) == "Q" {
			return false
		}
	}

	i, j := row - 1, col + 1
	for i >= 0 && j < len(board) {
		if string(board[i][j]) == "Q" {
			return false
		}
		i--
		j++
	}

	i, j = row - 1, col -1
	for i >= 0 && j >= 0 {
		if string(board[i][j]) == "Q" {
			return false
		}
		i--
		j--
	}

	return true
}
