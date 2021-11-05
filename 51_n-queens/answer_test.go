/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2021/6/7 10:26 下午
 */

package _n_queens

import (
	"reflect"
	"testing"
)

func TestAnswer(t *testing.T) {
	tests := []struct {
		input  int
		expect [][]string
	}{
		{
			4,
			[][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for i, test := range tests {
		if actual := solveNQueens(test.input); !reflect.DeepEqual(test.expect, actual) {
			t.Errorf("%d: input = %d, expect = %+v, actual = %+v", i, test.input, test.expect, actual)
		}
	}
}
