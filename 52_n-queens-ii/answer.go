/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/n-queens-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/28 17:24
 */

package _n_queens_ii

func totalNQueens(n int) int {
	var cnt int
	var backtrack func([][]int, int)

	backtrack = func(board [][]int, row int) {
		if row == len(board) {
			cnt++
			return
		}

		for col := 0; col < n; col++ {
			if !isValid(board, row, col) {
				continue
			}
			board[row][col] = 1
			backtrack(board, row+1)
			board[row][col] = 0
		}
	}

	var board [][]int
	for i := 0; i < n; i++ {
		board = append(board, make([]int, n))
	}

	backtrack(board, 0)

	return cnt
}

func isValid(board [][]int, row, col int) bool {
	for i := 0; i < len(board); i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	i, j := row-1, col+1
	for i >= 0 && j < len(board) {
		if board[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	i, j = row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	return true
}
