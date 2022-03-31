/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/diagonal-traverse-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/31 17:18
 */

package _diagonal_traverse_ii

func findDiagonalOrder(nums [][]int) []int {
	var ans []int

	var rec [][]int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i+j < len(rec) {
				rec[i+j] = append(rec[i+j], nums[i][j])
			} else {
				rec = append(rec, []int{nums[i][j]})
			}
		}
	}

	for i := 0; i < len(rec); i++ {
		for j := len(rec[i]) - 1; j >= 0; j-- {
			ans = append(ans, rec[i][j])
		}
	}

	return ans
}
