/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/projection-area-of-3d-shapes/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/4/26 12:59
 */

package _projection_area_of_3d_shapes

func projectionArea(grid [][]int) int {
	var xy, xz, yz int
	for i, row := range grid {
		var xHeight, yHeihgt int
		for j, num := range row {
			if num > 0 {
				xy++
			}
			xHeight = max(xHeight, num)
			yHeihgt = max(yHeihgt, grid[j][i])
		}
		xz += xHeight
		yz += yHeihgt
	}
	return xy + xz + yz
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
