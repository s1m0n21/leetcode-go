/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/pascals-triangle-ii/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/2/13 4:11 PM
 */

package _pascals_triangle_ii

func getRow(rowIndex int) []int {
	var ans = []int{1}
	for i := 1; i < rowIndex+1; i++ {
		t := []int{1}
		for j := 1; j < i; j++ {
			t = append(t, ans[j-1]+ans[j])
		}
		t = append(t, 1)
		ans = t
	}
	return ans
}
