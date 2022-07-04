/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/minimum-absolute-difference/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/7/4 09:29
 */

package _minimum_absolute_difference

import "sort"

func minimumAbsDifference(arr []int) [][]int {
	var min int
	var ret [][]int

	sort.Ints(arr)
	for i := 1; i < len(arr); i++ {
		n := arr[i] - arr[i-1]
		if n < min || min == 0 {
			min = n
			ret = nil
		}

		if n == min {
			ret = append(ret, []int{arr[i-1], arr[i]})
		}
	}

	return ret
}
