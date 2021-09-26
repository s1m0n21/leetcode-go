/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/word-search/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/9/26 1:21 下午
 */

package _word_search

var offset = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var m, n int
var visited [][]bool

func exist(board [][]byte, word string) bool {
	m, n = len(board), len(board[0])
	visited = nil
	for i := 0; i < m; i++ {
		visited = append(visited, make([]bool, n))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if search(board, word, 0, i, j) {
				return true
			}
		}
	}

	return false
}

func search(board [][]byte, word string, idx, x, y int) bool {
	if idx == len(word)-1 {
		return board[x][y] == word[idx]
	}

	if board[x][y] == word[idx] {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			nextx := x + offset[i][0]
			nexty := y + offset[i][1]
			if (nextx >= 0 && nextx < m && nexty >= 0 && nexty < n) &&
				(!visited[nextx][nexty]) &&
				(search(board, word, idx+1, nextx, nexty)) {
				return true
			}
		}
		visited[x][y] = false
	}

	return false
}
