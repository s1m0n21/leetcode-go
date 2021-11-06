/*
 * @Author       : s1m0n21
 * @Description  : Test answer
 * @Date         : 2021/11/05 4:54 PM
 */

package _sudoku_solver

import (
	"fmt"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input [][]byte
		//expect [][]byte
	}{
		{
			[][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			//[][]byte{
			//	{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
			//	{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
			//	{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
			//	{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
			//	{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
			//	{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
			//	{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
			//	{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
			//	{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
			//},
		},
	}

	for _, test := range tests {
		solveSudoku(test.input)
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				n := test.input[i][j]
				if !checkValid(test.input, i, j) {
					n = 'f'
				}

				if j == 8 {
					fmt.Printf("%s\n", string(n))
				} else {
					fmt.Printf("%s ", string(n))
				}
			}
		}
	}
}

func checkValid(board [][]byte, x, y int) bool {
	n := board[x][y]
	for i := 0; i < 9; i++ {
		if i == y {
			continue
		}
		if board[x][i] == n {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if i == x {
			continue
		}
		if board[i][y] == n {
			return false
		}
	}

	startx := (x / 3) * 3
	starty := (y / 3) * 3
	for i := startx; i < startx+3; i++ {
		for j := starty; j < starty+3; j++ {
			if i == x && j == y {
				continue
			}

			if board[i][j] == n {
				return false
			}
		}
	}

	return true
}
