/**
 * @Author         : s1m0n21
 * @Description    : Answer of https://leetcode-cn.com/problems/excel-sheet-column-title
 * @Project        : leetcode-go
 * @File           : answer.go
 * @Date           : 2022/3/14 11:31 AM
 */

package _excel_sheet_column_title

func convertToTitle(columnNumber int) string {
	var ans []byte

	for columnNumber > 0 {
		columnNumber--
		ans = append(ans, byte('A'+columnNumber%26))
		columnNumber /= 26
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}

	return string(ans)
}
