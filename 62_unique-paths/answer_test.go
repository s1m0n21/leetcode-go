/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/1/7 12:28 PM
 */

package _unique_paths

import (
	"testing"
)

func TestAnswer(t *testing.T) {
	type input struct {
		m, n int
	}

	tests := []struct {
		input  input
		expect int
	}{
		{input{3, 7}, 28},
		{input{3, 2}, 3},
		{input{7, 3}, 28},
		{input{3, 3}, 6},
	}

	for i, test := range tests {
		if actual := uniquePaths(test.input.m, test.input.n); actual != test.expect {
			t.Errorf("%d: input = %+v, exepct = %d, actual = %d", i, test.input, test.expect, actual)
		}
	}
}
