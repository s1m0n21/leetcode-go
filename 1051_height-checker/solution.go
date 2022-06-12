/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/height-checker/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/6/13 00:57
 */

package _height_checker

import "sort"

func heightChecker(heights []int) int {
	sorted := append([]int(nil), heights...)
	sort.Ints(sorted)

	var diff int
	for i, h := range heights {
		if sorted[i] != h {
			diff++
		}
	}

	return diff
}
