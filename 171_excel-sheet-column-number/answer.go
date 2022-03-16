/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/excel-sheet-column-number/
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/16 11:34 AM
 */

package _excel_sheet_column_number

func titleToNumber(columnTitle string) int {
	var ans int
	for i, m := len(columnTitle)-1, 1; i >= 0; i-- {
		n := int(columnTitle[i]-'A') + 1
		ans += n * m
		m *= 26
	}
	return ans
}
