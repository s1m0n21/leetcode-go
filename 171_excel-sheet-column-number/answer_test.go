/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/3/16 11:42 AM
 */

package _excel_sheet_column_number

import "testing"

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
		{"FXSHRXW", 2147483647},
	}

	for i, test := range tests {
		if actual := titleToNumber(test.input); actual != test.expect {
			t.Errorf("%d: input = %s, expect = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
