/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/find-all-duplicates-in-an-array/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/8 01:23
 */

package _find_all_duplicates_in_an_array

func findDuplicates(nums []int) []int {
	var ret []int

	s := make([]int, len(nums)+1)
	for _, n := range nums {
		s[n]++
		if s[n] == 2 {
			ret = append(ret, n)
		}
	}

	return ret
}
