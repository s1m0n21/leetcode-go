/**
 * @Author         : s1m0n21
 * @Description    : Solution of https://leetcode.cn/problems/delete-columns-to-make-sorted/
 * @Project        : leetcode-go
 * @File           : solution.go
 * @Date           : 2022/5/12 09:47
 */

package _delete_columns_to_make_sorted

func minDeletionSize(strs []string) int {
	var ret int

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				ret++
				break
			}
		}
	}

	return ret
}
