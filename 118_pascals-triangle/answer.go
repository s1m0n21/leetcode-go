/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/pascals-triangle/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2021/12/27 12:46 PM
 */

package _pascals_triangle

func generate(numRows int) [][]int {
	var f = [][]int{{1}}
	for i := 1; i < numRows; i++ {
		f = append(f, []int{1})
		for j := 1; j < i; j++ {
			f[i] = append(f[i], f[i-1][j-1]+f[i-1][j])
		}
		f[i] = append(f[i], 1)
	}
	return f
}
