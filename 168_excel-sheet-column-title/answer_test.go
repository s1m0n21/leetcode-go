/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/14 11:37 AM
 */

package _excel_sheet_column_title

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
		{2147483647, "FXSHRXW"},
	}

	for i, test := range tests {
		if actual := convertToTitle(test.input); actual != test.expect {
			t.Errorf("%d: input = %d, expect = %s, actual = %s", i, test.input, test.expect, actual)
		}
	}
}
