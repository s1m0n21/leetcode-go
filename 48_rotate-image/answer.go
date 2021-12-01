/**
 * @Author         : s1m0n21
 * @Description    : Answer for https://leetcode-cn.com/problems/rotate-image/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/11/30 10:07 上午
 */

package _rotate_image

func rotate(matrix [][]int) {
	var stack []int
	n := len(matrix)
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			stack = append(stack, matrix[j][i])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			m := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			matrix[i][j] = m
		}
	}
}
