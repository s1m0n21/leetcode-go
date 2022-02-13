/**
 * @Author         : s1m0n21
 * @Description    : Test answer
 * @Project        : leetcode-go
 * @File           : answer_test.go
 * @Date           : 2022/2/13 10:10 PM
 */

package _search_a_2d_matrix_ii

import "testing"

func TestAnswer(t *testing.T) {
	type input struct {
		matrix [][]int
		target int
	}

	tests := []struct {
		input  input
		expect bool
	}{
		{input{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5}, true},
		{input{[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20}, false},
	}

	for i, test := range tests {
		if actual := searchMatrix(test.input.matrix, test.input.target); actual != test.expect {
			t.Errorf("%d: input = %+v; expect = %t, actual = %t", i, test.input, test.expect, actual)
		}
	}
}
